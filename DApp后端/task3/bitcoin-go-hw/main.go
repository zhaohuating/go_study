package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main1() {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		//Host:         "bitcoin-rpc.publicnode.com:443",
		Host:         "bitcoin-mainnet.publicnode.com:443",
		User:         "",
		Pass:         "",
		HTTPPostMode: true,
		DisableTLS:   true,
		// 关键：显式指定 JSON-RPC 1.0 头 + UA，避免 CDN 拦截
		ExtraHeaders: map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   "btcd/0.23",
			"Accept":       "application/json",
		},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	height, err := client.GetBlockCount()
	if err != nil {
		log.Fatalf("GetBlockCount error: %v", err)
	}
	log.Printf("当前区块高度 %d", height)
}

func main() {

	// 配置RPC连接参数，精确匹配curl命令
	connCfg := &rpcclient.ConnConfig{
		Host:         "bitcoin-mainnet.core.chainstack.com:443",
		User:         "jovial-varahamihira",                    //
		Pass:         "audio-tartar-turf-cactus-deacon-speech", //
		HTTPPostMode: true,                                     //
		DisableTLS:   false,
	}

	// 创建RPC客户端
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatalf("创建RPC客户端失败: %v", err)
	}
	defer client.Shutdown()

	// 调用getblockcount方法获取区块高度
	blockHeight, err := client.GetBlockCount()
	if err != nil {
		log.Fatalf("获取区块高度失败: %v", err)
	}
	blockChainInfo, err := client.GetBlockChainInfo()

	if err != nil {
		log.Fatalf("获取区块链信息失败: %v", err)
	}

	// 输出结果
	fmt.Printf("当前比特币区块高度: %d\n", blockHeight)
	fmt.Printf("当前比特币区块链信息: %d\n", blockChainInfo)
}
