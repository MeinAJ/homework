package service

import (
	"context"
	"crypto/ecdsa"
	"dapp/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func BeggingContractService() {
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
	contractAddress := common.HexToAddress(ContractAddr)
	instance, err := contract.NewBeggingContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// 执行捐款 (0.0001 ETH)
	amount := big.NewInt(100000000000000) // 0.0001 ETH in wei
	auth.Value = amount
	tx, err := instance.Donate(auth)
	if err != nil {
		log.Fatalf("Failed to donate: %v", err)
	}
	fmt.Printf("Donation sent! Tx hash: %s\n", tx.Hash().Hex())

	// 查询总捐款额
	total, err := instance.TotalDonations(nil)
	if err != nil {
		log.Fatalf("Failed to get total donations: %v", err)
	}
	fmt.Printf("Total donations: %s wei\n", total.String())

	// 查询当前账户捐款额
	userDonation, err := instance.GetDonation(nil, fromAddress)
	if err != nil {
		log.Fatalf("Failed to get user contract: %v", err)
	}
	fmt.Printf("Your contract: %s wei\n", userDonation.String())

	// 提取捐款
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth.Value = big.NewInt(0) // 非 payable 函数设为 0
	withdraw, err := instance.Withdraw(auth)
	if err != nil {
		log.Fatalf("Failed to get withdraw: %v", err)
	}
	fmt.Printf("Withdraw sent! Tx hash: %s\n", withdraw.Hash().Hex())
}
