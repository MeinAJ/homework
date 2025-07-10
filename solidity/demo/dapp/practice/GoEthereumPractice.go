package practice

import (
	"context"
	"crypto/ecdsa"
	"dapp/contract"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	SepoliaURL   = os.Getenv("SEPOLIA_URL")
	PrivateKey   = os.Getenv("WALLET_PRIVATE_KEY")
	ContractAddr = os.Getenv("CONTRACT_ADDRESS")
	Address      = os.Getenv("WALLET_ADDRESS")
)

// GetHeaderByNumber 获取区块头信息
func GetHeaderByNumber(client *ethclient.Client, background context.Context, number *big.Int) {
	header, err := client.HeaderByNumber(background, number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("区块头信息：%v\n", header)
}

func TransactionToMessage(client *ethclient.Client, tx *types.Transaction, chainID *big.Int) (*core.Message, error) {
	signer := types.LatestSignerForChainID(chainID)
	var baseFee *big.Int
	if tx.Type() == types.DynamicFeeTxType {
		// 实际应用中应从节点获取最新 baseFee
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		baseFee = header.BaseFee
	}
	return core.TransactionToMessage(tx, signer, baseFee)
}

// GetBlockTransactionsByNumber 获取区块交易信息
func GetBlockTransactionsByNumber(client *ethclient.Client, background context.Context, number *big.Int) {
	block, err := client.BlockByNumber(background, number)
	if err != nil {
		log.Fatal(err)
	}
	for _, tx := range block.Transactions() {
		// 打印每个交易的各项信息
		fmt.Printf("交易hash：%s\n", tx.Hash().Hex())
		fmt.Printf("交易nonce：%d\n", tx.Nonce())
		fmt.Printf("交易gasPrice：%d\n", tx.GasPrice())
		fmt.Printf("交易gasLimit：%d\n", tx.Gas())
		fmt.Printf("交易value：%d\n", tx.Value())
		message, err := TransactionToMessage(client, tx, nil)
		if err != nil {
			fmt.Println("获取发送方地址失败：", err)
		}
		fmt.Printf("交易from：%s\n", message.From.Hex())
		fmt.Printf("交易to：%s\n", tx.To().Hex())
		fmt.Printf("交易input：%s\n", hexutil.Encode(tx.Data()))
	}
}

// GetBlockTransactionsByHashByNumber 获取区块交易信息
func GetBlockTransactionsByHashByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) {
	block, err := client.BlockByNumber(ctx, number)
	if err != nil {
		log.Fatal(err)
	}
	blockHash := common.HexToHash(block.Hash().Hex())
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			fmt.Println("获取交易失败：", err)
		}
		// 打印每个交易的各项信息
		fmt.Printf("交易hash：%s\n", tx.Hash().Hex())
		fmt.Printf("交易nonce：%d\n", tx.Nonce())
		fmt.Printf("交易gasPrice：%d\n", tx.GasPrice())
		fmt.Printf("交易gasLimit：%d\n", tx.Gas())
		fmt.Printf("交易value：%d\n", tx.Value())
		message, err := TransactionToMessage(client, tx, nil)
		if err != nil {
			fmt.Println("获取发送方地址失败：", err)
		}
		fmt.Printf("交易from：%s\n", message.From.Hex())
		fmt.Printf("交易to：%s\n", tx.To().Hex())
		fmt.Printf("交易input：%s\n", hexutil.Encode(tx.Data()))
	}
}

// GetBlockTransactionBySingleHashByNumber 获取区块交易信息
func GetBlockTransactionBySingleHashByNumber(client *ethclient.Client, background context.Context, number *big.Int) {
	txHash := common.HexToHash("0x373ec1d3d0319ddaff9a960df73ad118798f49f8a0454575ffb26de20f1942f1")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println("获取交易失败：", err)
	}
	// 打印每个交易的各项信息
	fmt.Printf("交易hash：%s\n", tx.Hash().Hex())
	fmt.Printf("交易nonce：%d\n", tx.Nonce())
	fmt.Printf("交易gasPrice：%d\n", tx.GasPrice())
	fmt.Printf("交易gasLimit：%d\n", tx.Gas())
	fmt.Printf("交易value：%d\n", tx.Value())
	message, err := TransactionToMessage(client, tx, nil)
	if err != nil {
		fmt.Println("获取发送方地址失败：", err)
	}
	fmt.Printf("交易from：%s\n", message.From.Hex())
	fmt.Printf("交易to：%s\n", tx.To().Hex())
	fmt.Printf("交易input：%s\n", hexutil.Encode(tx.Data()))
}

// GetReceiptsByHashByNumber 获取区块交易信息
func GetReceiptsByHashByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) {
	block, err := client.BlockByNumber(ctx, number)
	if err != nil {
		log.Fatal(err)
	}
	blockHash := common.HexToHash(block.Hash().Hex())
	receiptByHash, err := client.BlockReceipts(ctx, rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	receiptsByNum, err := client.BlockReceipts(ctx, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(block.Number().Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true
	for _, receipt := range receiptByHash {
		fmt.Printf("交易hash：%s\n", receipt.TxHash.Hex())
		fmt.Printf("交易状态：%v\n", receipt.Status)
		fmt.Printf("交易gasUsed：%d\n", receipt.GasUsed)
		fmt.Printf("交易logs：%v\n", receipt.Logs)
		fmt.Printf("交易TransactionIndex：%v\n", receipt.TransactionIndex)
		fmt.Printf("交易ContractAddress：%v\n", receipt.ContractAddress.Hex())
	}
}

func CreateNewWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥（去掉前2位）：", hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥（去掉前4位）：", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	// 第一种方式：通过公钥计算地址，使用crypto.PubkeyToAddress
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("第一种方式，钱包地址：", address)
	// 第二种方式：通过公钥的Keccak256哈希值计算地址，取后20个字节（40个字符）作为地址
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("第二种方式，公钥的Keccak256哈希值：", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("第二种方式，公钥的Keccak256哈希值的后40个字符（钱包地址）：", hexutil.Encode(hash.Sum(nil)[12:]))
}

func EthTransfer(client *ethclient.Client, ctx context.Context) {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (0.001 eth)
	gasLimit := uint64(21000)             // in units
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(Address)
	var data []byte
	tx := types.NewTx(&types.LegacyTx{Nonce: nonce, To: &toAddress, Value: value, Gas: gasLimit, GasPrice: gasPrice, Data: data})

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func TokenTransfer(client *ethclient.Client, ctx context.Context) {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("将token发送给谁")
	tokenAddress := common.HexToAddress("合约地址")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	amount := new(big.Int)
	amount.SetString("100000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256
	tx := types.NewTx(&types.LegacyTx{Nonce: nonce, To: &tokenAddress, Value: value, Gas: gasLimit, GasPrice: gasPrice, Data: data})

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func GetBalance(client *ethclient.Client, ctx context.Context) {
	account := common.HexToAddress("0xb052360268d7F2FA4Eb62eD1c6194257935C0BfE")
	balanceAt, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("最新区块的账户余额（以wei为单位）：", balanceAt)
	blockNumber := big.NewInt(8475327)
	balanceAt, err = client.BalanceAt(ctx, account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块高度为8475327的账户余额（以wei为单位）：", balanceAt) // 25729324269165216042
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("区块高度为8475327的账户余额（以eth为单位）：", ethValue) // 25.729324269165216041
	pendingBalance, err := client.PendingBalanceAt(ctx, account)
	fmt.Println("区块高度为8475327的账户未确认的余额（以wei为单位）：", pendingBalance) // 25729324269165216042
}

func CheckAndSubscribeNewBlock(wsUrl string, ctx context.Context) {
	wsClient, err := ethclient.DialContext(ctx, wsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer wsClient.Close()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 添加安全退出机制
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		cancel()
	}()

	SubscribeNewBlock(wsClient, ctx)
}

func SubscribeNewBlock(wsClient *ethclient.Client, ctx context.Context) {
	headers := make(chan *types.Header, 128) // 增加缓冲区大小
	sub, err := wsClient.SubscribeNewHead(ctx, headers)
	if err != nil {
		log.Fatal("subscribe new head error:", err)
	}
	defer sub.Unsubscribe()

	// 用于跟踪最近区块
	recentBlocks := make(map[common.Hash]bool)

	for {
		select {
		case <-ctx.Done():
			log.Println("上下文取消，退出订阅")
			return

		case err := <-sub.Err():
			log.Printf("订阅错误: %v", err)
			// 尝试重新订阅
			time.Sleep(5 * time.Second)
			sub, err = wsClient.SubscribeNewHead(ctx, headers)
			if err != nil {
				log.Fatal("重新订阅失败:", err)
			}

		case header := <-headers:
			hash := header.Hash()

			// 检查是否已处理过此区块
			if recentBlocks[hash] {
				continue
			}
			recentBlocks[hash] = true

			// 清理旧区块
			if len(recentBlocks) > 100 {
				for h := range recentBlocks {
					delete(recentBlocks, h)
					break
				}
			}

			log.Printf("收到新区块头 #%d: %s", header.Number, hash.Hex())

			// 使用重试机制获取区块
			block, err := getBlockWithRetry(wsClient, ctx, hash, 3, 500*time.Millisecond)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					log.Printf("区块可能已被丢弃 (分叉): %s", hash.Hex())
				} else {
					log.Printf("获取区块失败: %v", err)
				}
				continue
			}

			// 验证区块头哈希是否匹配
			if block.Hash() != hash {
				log.Printf("警告: 哈希不匹配! 头: %s, 区块: %s",
					hash.Hex(), block.Hash().Hex())
			}

			// 处理区块
			processBlock(block)
		}
	}
}

func getBlockWithRetry(client *ethclient.Client, ctx context.Context,
	hash common.Hash, retries int, delay time.Duration) (*types.Block, error) {
	for i := 0; i < retries; i++ {
		block, err := client.BlockByHash(ctx, hash)
		if err == nil {
			return block, nil
		}

		// 如果是"not found"错误，可能是区块传播延迟
		if strings.Contains(err.Error(), "not found") {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(delay):
				// 等待后重试
			}
			continue
		}

		return nil, err
	}
	return nil, fmt.Errorf("重试 %d 次后仍无法获取区块", retries)
}

func processBlock(block *types.Block) {
	fmt.Println("区块哈希:", block.Hash().Hex())
	fmt.Println("区块高度:", block.Number().Uint64())
	fmt.Println("区块时间:", block.Time())
	fmt.Println("交易数量:", len(block.Transactions()))
}

func DeployContract(client *ethclient.Client, ctx context.Context) {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatal("get private key error:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal("get pending nonce error:", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal("get suggested gas price error:", err)
	}

	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal("get chain id error:", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal("get transactor error:", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := contract.DeployMyToken(auth, client)
	if err != nil {
		log.Fatal("deploy contract error:", err)
	}

	fmt.Println("contract address:", address.Hex())
	fmt.Println("tx hash:", tx.Hash().Hex())

	_ = instance
}
