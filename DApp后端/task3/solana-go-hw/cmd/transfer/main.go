package main

import (
	"context"
	"fmt"
	"log"
	rpc2 "task3/solana-go-hw/rpc"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

// --------- 辅助：指针工具 ----------
func boolPtr(b bool) *bool { return &b }
func uintPtr(u uint) *uint { return &u }

// --------- 常量 ----------
const (
	// 私钥文件 & 收款地址
	keyFile = "/Users/zhaohuating/.config/solana/testid1.json"
	toAddr  = "ChD6SjzQv9T8UvU2ERxCC7bQ1egcc1S27dTcdUuvf1jo"
	// 转账金额
	amountSol = 0.0005
)

// --------- 主流程 ----------
func main() {
	transferAndListen()
}

func main1() {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := ws.Connect(context.Background(), rpc2.TestNet_WS)
	if err != nil {
		panic(err)
	}
	defer cancle()
	log.Println("wss 已建立 ...")

	sub, err := client.SlotSubscribe()
	if err != nil {
		panic(err)
	}

	for {
		got, err := sub.Recv(ctx)
		if err != nil {
			panic(err)
		}
		spew.Dump(got)
	}
}

// 统一错误包装
func check(err error) {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}

// 转账 + 监听
func transferAndListen() {
	ctx := context.Background()

	// 1. 连 RPC & WebSocket
	client := rpc.New(rpc2.QuicknodeTestNet_RPC)
	wsClient, err := ws.Connect(ctx, rpc2.QuicknodeTestNet_WS)
	check(err)
	defer wsClient.Close()

	// 2. 加载私钥
	sender, err := solana.PrivateKeyFromSolanaKeygenFile(keyFile)
	check(err)
	fromPub := sender.PublicKey()
	toPub := solana.MustPublicKeyFromBase58(toAddr)

	// 3. 拿最新 blockhash
	recent, err := client.GetLatestBlockhash(ctx, rpc.CommitmentConfirmed)
	check(err)

	// 4. 构造交易
	lamports := uint64(amountSol * 1e9)
	ix := system.NewTransferInstruction(lamports, fromPub, toPub).Build()
	tx, err := solana.NewTransaction(
		[]solana.Instruction{ix},
		recent.Value.Blockhash,
		solana.TransactionPayer(fromPub),
	)
	check(err)

	// 5. 签名
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key == fromPub {
			return &sender
		}
		return nil
	})
	check(err)

	// 6. 提前订阅签名
	sig := tx.Signatures[0]
	sub, err := wsClient.SignatureSubscribe(solana.MustSignatureFromBase58(sig.String()), rpc.CommitmentConfirmed)
	fmt.Println("已订阅签名：", sig)
	check(err)
	defer sub.Unsubscribe()

	// 7. 启动 goroutine 等推送
	doneCh := make(chan struct{}, 1)
	go func() {
		fmt.Println("等待订阅消息...")
		for {
			select {
			case got := <-sub.Response():
				fmt.Printf(">>> 交易确认通知：slot=%d, err=%v\n", got.Context.Slot, got.Value.Err)
				doneCh <- struct{}{}
			case err := <-sub.Err():
				log.Printf("订阅错误: %v", err)
				return
			}
		}
	}()

	// 8. 睡 500ms 确保订阅报文到达节点
	time.Sleep(500 * time.Millisecond)

	// 9. 发送交易（带重试）
	fmt.Printf("正在发送交易 %s ...\n", sig)
	_, err = client.SendTransactionWithOpts(ctx, tx,
		rpc.TransactionOpts{
			SkipPreflight:       false,
			MaxRetries:          uintPtr(5),
			PreflightCommitment: rpc.CommitmentConfirmed,
		},
	)
	check(err)
	// 10. 等确认或超时
	select {
	case <-doneCh:
		fmt.Println("✅ 完成")
	case <-time.After(60 * time.Second):
		fmt.Println("⏰ 超时未收到确认")
	}
}
