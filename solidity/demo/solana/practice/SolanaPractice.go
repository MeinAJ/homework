package practice

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"os"
)

var AlchemyUrl string

func init() {
	if os.Getenv("ALCHEMY_URL") == "" {
		AlchemyUrl = rpc.DevnetRPCEndpoint
	} else {
		AlchemyUrl = os.Getenv("ALCHEMY_URL")
	}
}

// QueryBlockInfo 查询区块信息
func QueryBlockInfo() {
	// 创建RPC客户端
	solanaClient := client.NewClient(AlchemyUrl)
	// 获取最新区块高度
	slot, err := solanaClient.GetSlot(context.Background())
	if err != nil {
		panic("获取区块高度失败: " + err.Error())
	}
	// 获取最新区块
	latestBlock, err := solanaClient.GetBlock(context.Background(), slot)
	if err != nil {
		panic("查询区块失败: " + err.Error())
	}
	fmt.Printf("最新区块高度: %d\n", slot)
	fmt.Printf("区块哈希: %s\n", latestBlock.Blockhash)
	fmt.Printf("区块高度: %d\n", latestBlock.BlockHeight)
	fmt.Printf("交易数量: %d\n", len(latestBlock.Transactions))
}

// QuerySolanaVersion 查询solana版本
func QuerySolanaVersion() {
	solanaClient := client.NewClient(AlchemyUrl)
	version, _ := solanaClient.GetVersion(context.Background())
	fmt.Println("Solana node version:", version.SolanaCore)
}
