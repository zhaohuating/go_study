package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	rpc2 "task3/solana-go-hw/rpc"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
	// 1. 连接 DevNet WebSocket
	wsClient, err := ws.Connect(context.Background(), rpc2.TestNet_WS)
	if err != nil {
		log.Fatalf("WebSocket 连接失败: %v", err)
	}
	defer wsClient.Close()

	// 2. 演示：订阅某笔交易签名（可换成自己刚发出的）
	txSig := solana.MustSignatureFromBase58("52nYh4ABNYB3nZBBF9yC9m93M4k96mzAAjXYj93x5MBKzErUFVCeAdztfsS95WABctGz7wG2qwjypV3HCxxS7Ct")
	sub, err := wsClient.SignatureSubscribe(txSig, rpc.CommitmentConfirmed)
	if err != nil {
		log.Fatalf("订阅交易失败: %v", err)
	}
	defer sub.Unsubscribe()

	// 3. 打印推送结果
	go func() {
		for {
			select {
			case got := <-sub.Response():
				fmt.Printf(">>> 交易确认通知：slot=%d, err=%v\n", got.Context.Slot, got.Value.Err)
			case err := <-sub.Err():
				log.Printf("订阅错误: %v", err)
				return
			}
		}
	}()

	// 4. 优雅退出
	fmt.Println("监听中，按 Ctrl-C 退出...")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("已退出监听")
}
