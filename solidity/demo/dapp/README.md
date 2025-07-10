# 初始化项目

go mod init dapp

# 安装依赖包

go mod tidy

# 安装go-ethereum

go get -u github.com/ethereum/go-ethereum

# 安装abigen

go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# 编译合约

abigen --abi artifacts/contracts/BeggingContract.sol/BeggingContract.json --pkg BeggingContract --out BeggingContract.go

# 编译go文件

go build -o dapp main.go

# 将abi文件和bin文件放入同一目录下，使用abigen生成go文件

# 先使用npx hardhat compile生成abi和bin文件
npx hardhat compile

# 再使用abigen生成go文件
abigen --abi=dapp/sol_source/BeggingContract.abi.json --bin=dapp/sol_source/BeggingContract.bin --pkg=contract --type=BeggingContract --out=dapp/contract/BeggingContract.go

abigen --abi=dapp/sol_source/MyToken.abi.json --bin=dapp/sol_source/MyToken.bin --pkg=contract --type=MyToken --out=dapp/contract/MyToken.go
