package main

import (
	"context"
	"dapp/practice"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

var client, _ = ethclient.Dial(os.Getenv("SEPOLIA_URL"))
var wsUrl = os.Getenv("SEPOLIA_WS_URL")
var ctx = context.Background()

func main() {
	// MyToken 合约 & 查询代币
	//service.MyTokenService()
	// 部署 MyToken 合约
	// 捐款合约 & 查询捐款
	//service.BeggingContractDonate()
	// 查询区块头
	//practice.GetHeaderByNumber(client, ctx, nil)
	// 查询区块的所有交易（事物）
	//practice.GetBlockTransactionsByNumber(client, ctx, nil)
	// 查询区块的所有交易（事物数量）
	//practice.GetBlockTransactionsByHashByNumber(client, ctx, nil)
	// 查询单个交易（事物）
	//practice.GetBlockTransactionBySingleHashByNumber(client, ctx, nil)
	// 获取所有收据列表
	//practice.GetReceiptsByHashByNumber(client, ctx, nil)
	// 创建新钱包
	//practice.CreateNewWallet()
	// ETH转账
	//practice.EthTransfer(client, ctx)
	// 代币转账
	//practice.TokenTransfer(client, ctx)
	// 查询账户余额
	//practice.GetBalance(client, ctx)
	// 订阅最新区块
	//practice.CheckAndSubscribeNewBlock(wsUrl, ctx)
	// 部署合约
	//practice.DeployContract(client, ctx)
	// 调用合约
	practice.CallContract(client, ctx)
}
