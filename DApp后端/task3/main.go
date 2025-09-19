package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

const (
	rpcEndpoint = "https://solana-testnet-rpc.publicnode.com"
	//pub         = "9LKpbP2pmsxVepFHBR1bZmibkQuwudgSzgS5YTKkuybL"
	pub = "Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS"
)

func main1() {
	//cluster := rpc.MainNetBeta
	//rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(cluster.RPC, rate.Every(time.Second), 5))

	//client := rpc.New(rpc.DevNet_RPC)
	client := rpc.New(rpcEndpoint)
	resp, err := client.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest blockhash:", resp.Value.Blockhash)

	getBalance(pub)

}

func main() {
	// 1. 连接公开节点
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         "bitcoin-rpc.publicnode.com:443", // 443 必须显式
		User:         "",                               // 公开节点无认证
		Pass:         "",
		HTTPPostMode: true,
		DisableTLS:   false, // 用 https，保持 false
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// 2. 查最新区块高度
	height, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("当前区块高度 %d", height)
}

//{"address":"bc1q695z03z6kweljcvpwft7vfu6kd0guf24yaaht2",
//"chain_stats":{"funded_txo_count":1149,"funded_txo_sum":347426115611,"spent_txo_count":651,"spent_txo_sum":202508630677,"tx_count":1150},
//"mempool_stats":{"funded_txo_count":0,"funded_txo_sum":0,"spent_txo_count":0,"spent_txo_sum":0,"tx_count":0}}

type AddrInfo struct {
	Address    string `json:"address"`
	ChainStats struct {
		FundedTxoCount int   `json:"funded_txo_count"`
		FundedTxoSum   int64 `json:"funded_txo_sum"`
		SpentTxoCount  int   `json:"spent_txo_count"`
		SpentTxoSum    int64 `json:"spent_txo_sum"`
		TxCount        int   `json:"tx_count"`
	} `json:"chain_stats"`
	MempoolStats struct {
		FundedTxoCount int `json:"funded_txo_count"`
		FundedTxoSum   int `json:"funded_txo_sum"`
		SpentTxoCount  int `json:"spent_txo_count"`
		SpentTxoSum    int `json:"spent_txo_sum"`
		TxCount        int `json:"tx_count"`
	} `json:"mempool_stats"`
}

func main2() {
	addr := "bc1q695z03z6kweljcvpwft7vfu6kd0guf24yaaht2"
	url := fmt.Sprintf("https://blockstream.info/api/address/%s", addr)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var addrInfo AddrInfo
	if err = json.NewDecoder(resp.Body).Decode(&addrInfo); err != nil {
		panic(err)
	}

	confirmed := addrInfo.ChainStats.FundedTxoSum - addrInfo.ChainStats.SpentTxoSum
	pending := addrInfo.MempoolStats.FundedTxoSum - addrInfo.MempoolStats.SpentTxoSum
	total := confirmed + int64(pending)
	fmt.Printf("余额 %.8f BTC\n", float64(total)/1e8)
}

func getBalance(pub string) {
	client := rpc.New(rpcEndpoint)
	ctx := context.Background()
	publicKey := solana.MustPublicKeyFromBase58(pub)
	balance, err := client.GetBalance(ctx, publicKey, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance:", float64(balance.Value)/10e9)
}
