package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"task1/helper"
	"task1/store"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PRIVATEKEY   = "25d706fa02cdc95521bec870582c8b95e0e400784a6f948471044273ce31ef62"
	RAWURL       = "https://ethereum-sepolia-rpc.publicnode.com"
	WSSRAWURL    = "wss://ethereum-sepolia-rpc.publicnode.com"
	contractAddr = "0x65Aeb47eb7Eb153306BA87398544437D5560b162"
)

func main() {
	//BlockInfo(0)
	//TestTransferERC20()
	// SubscribeNewHead()
	//DeployContract()
	//LoadContract()
	ExeContract()

}

func BlockInfo(blockNum int64) {
	client, err := ethclient.Dial(RAWURL)

	if err != nil {
		helper.HandleErr("Dial", err)
	}

	newInt := big.NewInt(blockNum)
	if blockNum == 0 {
		newInt = nil
	}

	block, err := client.BlockByNumber(context.Background(), newInt)
	if err != nil {
		helper.HandleErr("BlockByNumber", err)
	}

	blockNumber := block.Number()
	fmt.Println("区块高度：", blockNumber)
	hash := block.Hash().Hex()
	fmt.Println("区块hash：", hash)

	time := block.Time()
	fmt.Println("交易时间：", helper.FormatTimestampCN(int64(time)))
	transactionCount := len(block.Transactions())
	fmt.Println("交易数量：", transactionCount)
}

func TestTransferERC20() {
	// 连接测试网
	client, err := ethclient.Dial(RAWURL)
	if err != nil {
		helper.HandleErr("Dial", err)
	}

	//加载私钥
	privateKey, err := crypto.HexToECDSA(PRIVATEKEY)
	if err != nil {
		helper.HandleErr("HexToECDSA", err)
	}
	from := privateKey.Public().(*ecdsa.PublicKey)
	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*from))
	if err != nil {
		helper.HandleErr("PendingNonceAt", err)
	}
	fmt.Println("nonce", nonce)

	// 计算gas费用
	tipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		helper.HandleErr("SuggestGasTipCap", err)
	}
	fmt.Println("tipCap", tipCap)
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		helper.HandleErr("HeaderByNumber", err)
	}
	gasFee := big.NewInt(0).Add(tipCap, header.BaseFee)
	fmt.Println("gasFee", gasFee)
	//gasFee = gasFee.Mul(gasFee, big.NewInt(110)).Div(gasFee, big.NewInt(100))
	fmt.Println("gasFee", gasFee)
	//
	value := big.NewInt(0)

	//构建交易数据（calldata）
	transferFnSignature := []byte("transfer(address,uint256)")
	methodID := crypto.Keccak256(transferFnSignature)[:4]
	tokenAddress := common.HexToAddress("0xd6AF33A43593665Bbca08FC910B8eedDe4bB76Ac")
	toAddress := common.HexToAddress("0x3F03a613E52070c03289F2b314B0F0852f83d7Dc")
	toAddressLeftPad := common.LeftPadBytes(toAddress.Bytes(), 32)
	amount, _ := new(big.Int).SetString("1000000000000000000", 10)
	amountLeftPad := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, toAddressLeftPad...)
	data = append(data, amountLeftPad...)

	// 估算 GasLimit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: crypto.PubkeyToAddress(*from),
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		helper.HandleErr("EstimateGas", err)
	}

	//构建 EIP-1559 交易
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		helper.HandleErr("ChainID", err)
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &tokenAddress,
		GasTipCap: tipCap,
		GasFeeCap: gasFee,
		Gas:       gasLimit,
		Value:     value,
		Data:      data,
	})

	// 签名
	signTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		helper.HandleErr("SignTx", err)
	}
	// 广播
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		helper.HandleErr("SendTransaction", err)
	} else {
		fmt.Println("send tx success")
		fmt.Println("tx hash", signTx.Hash())
	}
}

func SubscribeNewHead() {
	client, err := ethclient.Dial(WSSRAWURL)
	if err != nil {
		helper.HandleErr("Dial", err)
	}
	headers := make(chan *types.Header)
	head, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		helper.HandleErr("SubscribeNewHead", err)
	}

	for {
		select {
		case header := <-headers:
			fmt.Println("区块hash：", header.Hash())
			fmt.Println("区块高度：", header.Number.Uint64())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				helper.HandleErr("BlockByHash", err)
			}
			fmt.Println("区块Nonce", block.Nonce())
			fmt.Println("交易数量：", len(block.Transactions()))
			fmt.Println("生成时间：", helper.FormatTimestampCN(int64(header.Time)))

			fmt.Println("=======================================")
		case err := <-head.Err():
			helper.HandleErr("SubscribeNewHead", err)
		}
	}
}

func DeployContract() {
	client, err := ethclient.Dial(RAWURL)
	if err != nil {
		helper.HandleErr("Dial", err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATEKEY)
	if err != nil {
		helper.HandleErr("Parse", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		helper.HandleErr("PendingNonceAt", err)
	}
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	helper.HandleErr("SuggestGasPrice", err)
	//}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		helper.HandleErr("NetworkID", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		helper.HandleErr("NewKeyedTransactorWithChainID", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(int64(6248173))

	input := "1.0"

	address, transaction, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		helper.HandleErr("DeployStore", err)
	}

	fmt.Println("address:", address.Hex())
	fmt.Println("transaction:", transaction.Hash())
	_ = instance
}

func LoadContract() {
	client, err := ethclient.Dial(RAWURL)
	if err != nil {
		helper.HandleErr("Dial", err)
	}

	newStore, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		helper.HandleErr("NewStore", err)
	}
	fmt.Println(newStore)
}

func ExeContract() {
	client, err := ethclient.Dial(RAWURL)
	if err != nil {
		helper.HandleErr("Dial", err)
	}

	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		helper.HandleErr("NewStore", err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATEKEY)
	if err != nil {
		helper.HandleErr("NewPrivateKey", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		helper.HandleErr("NetworkID", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		helper.HandleErr("NewTransactor", err)
	}

	opts.GasLimit = 300000

	//var key [32]byte
	//var value [32]byte
	//copy(key[:], []byte("demo_save_key"))
	//copy(value[:], "demo_save_value11111")
	//tx, err := storeContract.SetItem(opts, key, value)
	//if err != nil {
	//	helper.HandleErr("SetItem", err)
	//}
	//fmt.Println("tx hash", tx.Hash())
	countTx, err := storeContract.CountItem(opts)
	if err != nil {
		helper.HandleErr("CountItem", err)
	}

	fmt.Println("count tx", countTx.Hash())
	time.Sleep(15 * time.Second)
	//_value, err := storeContract.Items(&bind.CallOpts{Context: context.Background()}, key)
	if err != nil {
		helper.HandleErr("Items", err)
	}

	//fmt.Println("is value saving in contract equals to origin value:", _value == value)
}

func buildTransactOpts() *bind.TransactOpts {
	client, err2 := ethclient.Dial(RAWURL)
	if err2 != nil {
		helper.HandleErr("Dial", err2)
	}
	privateKey, err := crypto.HexToECDSA(PRIVATEKEY)
	if err != nil {
		helper.HandleErr("HexToECDSA", err)
	}
	from := privateKey.Public().(*ecdsa.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*from))
	tipCap, err2 := client.SuggestGasTipCap(context.Background())
	if err2 != nil {
		helper.HandleErr("SuggestGasTipCap", err)
	}

	block, err2 := client.BlockByNumber(context.Background(), nil)
	if err2 != nil {
		helper.HandleErr("BlockByNumber", err)
	}

	baseFee := block.BaseFee()
	GasFeeCap := big.NewInt(0).Add(baseFee, tipCap)

	var data []byte
	transferFnSignature := []byte("CountItem()")
	methodID := crypto.Keccak256(transferFnSignature)[:4]
	toAddress := common.HexToAddress(contractAddr)
	data = append(data, methodID...)
	client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: crypto.PubkeyToAddress(*from),
		To:   &toAddress,
		Data: data,
	})

	chainID, err2 := client.NetworkID(context.Background())
	if err2 != nil {
		helper.HandleErr("NetworkID", err)
	}
	fmt.Println("chainID:", chainID)
	//transactOpts, err2 := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	return &bind.TransactOpts{
		From:      crypto.PubkeyToAddress(*from),
		Nonce:     big.NewInt(int64(nonce)),
		Context:   context.Background(),
		GasLimit:  300000,
		GasTipCap: tipCap,
		GasFeeCap: GasFeeCap,
	}
}
