package helper

import (
	"context"
	"errors"
	"log"
	"strings"
)

// 把原来裸 log.Fatal 换成这个
func HandleErr(tag string, err error) {
	if err == nil {
		return
	}
	chainID := 11155111
	// 1. 用 errors.Is 判断“底层根因”
	if errors.Is(err, context.DeadlineExceeded) {
		log.Fatalf("[%s] RPC 请求超时，请检查网络或换 Sepolia 节点", tag)
	}

	// 2. 用 strings.Contains 匹配常用字符串
	errStr := strings.ToLower(err.Error())

	switch {
	case strings.Contains(errStr, "insufficient funds"):
		log.Fatalf("[%s] 地址 ETH 余额不足，无法支付 gas，去水龙头领水: https://sepoliafaucet.com", tag)

	case strings.Contains(errStr, "nonce too low"):
		log.Fatalf("[%s] nonce 重复使用，请用 client.PendingNonceAt 重新获取最新 nonce", tag)

	case strings.Contains(errStr, "replacement transaction underpriced"):
		log.Fatalf("[%s] 之前已有相同 nonce 的交易且 gasPrice 更高，可提高 GasTipCap 或等待上链", tag)

	case strings.Contains(errStr, "execution reverted"):
		// 如果是合约 revert，往往带 revert reason
		if strings.Contains(errStr, "ERC20: transfer amount exceeds balance") {
			log.Fatalf("[%s] 代币余额不足，请确认你在该合约里拥有足够的代币", tag)
		}
		log.Fatalf("[%s] 合约执行 revert，原因：%s", tag, err.Error())

	case strings.Contains(errStr, "gas required exceeds allowance"):
		log.Fatalf("[%s] 当前 RPC 允许的 gas 上限 < 估算值，可手动调大 GasLimit 或换节点", tag)

	case strings.Contains(errStr, "invalid sender"):
		log.Fatalf("[%s] 链 ID 设置错误，当前 chainID=%d，请确认是否为 Sepolia", tag, chainID)

	default:
		// 兜底：把原始错误也打出来，方便复制谷歌
		log.Fatalf("[%s] 未预料的错误：%v", tag, err)
	}
}
