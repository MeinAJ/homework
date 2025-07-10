package service

import (
	"context"
	"crypto/ecdsa"
	"dapp/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func MyTokenService() {
	client, err := ethclient.Dial(SepoliaURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum network: %v", err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 获取账户地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	// 创建交易签名者
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// 创建合约实例
	ContractAddr = "0x889F7A8e98EAeD5A43F4418335320Eb6BbD7ade2"
	contractAddress := common.HexToAddress(ContractAddr)
	instance, err := contract.NewMyToken(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// 调用合约方法（直接购买0.01ETH）
	auth.Value = big.NewInt(10000000000000000)
	receive, err := instance.Receive(auth)
	if err != nil {
		log.Fatalf("Failed to call Receive method: %v", err)
	}
	log.Printf("Receive transaction hash: %v", receive.Hash().Hex())

	// 查询当前人的token余额
	balance, err := instance.GetTokens(nil, fromAddress)
	if err != nil {
		log.Fatalf("Failed to call BalanceOf method: %v", err)
	}
	log.Printf("Current balance of %v: %v", fromAddress.Hex(), balance.String())

	// 提取出当前人的ETH余额
	auth.Value = nil
	auth.From = fromAddress
	eth, err := instance.WithdrawETH(auth)
	if err != nil {
		log.Fatalf("Failed to call Withdraw method: %v", err)
	}
	log.Printf("Withdraw transaction hash: %v", eth.Hash().Hex())

	// 查询erc20合约的各项信息
	name, err := instance.Name(nil)
	if err != nil {
		log.Fatalf("Failed to call Name method: %v", err)
	}
	log.Printf("Token name: %v", name)
	symbol, err := instance.Symbol(nil)
	if err != nil {
		log.Fatalf("Failed to call Symbol method: %v", err)
	}
	log.Printf("Token symbol: %v", symbol)
	decimals, err := instance.Decimals(nil)
	if err != nil {
		log.Fatalf("Failed to call Decimals method: %v", err)
	}
	log.Printf("Token decimals: %v", decimals)
}
