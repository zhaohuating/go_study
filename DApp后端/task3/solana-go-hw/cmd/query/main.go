package main

import (
	"context"
	"fmt"
	"log"
	locrpc "task3/solana-go-hw/rpc"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	// 1. 连接 DevNet
	client := rpc.New(locrpc.QuicknodeTestNet_RPC)

	// 2. 获取最新区块 hash
	resp, err := client.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("获取最新区块失败: %v", err)
	}
	fmt.Printf("Latest blockhash: %s\n", resp.Value.Blockhash)

	// 3. 查询账户余额（示例地址可换成自己的）
	pub := solana.MustPublicKeyFromBase58("D7SHVCWN1q4QgZK1NgqXKJRmujH1di9FenWC81Fo1zC7")
	bal, err := client.GetBalance(context.TODO(), pub, rpc.CommitmentConfirmed)
	if err != nil {
		log.Fatalf("查询余额失败: %v", err)
	}
	fmt.Printf("账户 %s 余额: %d lamports (%.4f SOL)\n", pub, bal.Value, float64(bal.Value)/1e9)
}
