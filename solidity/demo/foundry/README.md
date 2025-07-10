## Foundry

**Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust.**

Foundry consists of:

- **Forge**: Ethereum testing framework (like Truffle, Hardhat and DappTools).
- **Cast**: Swiss army knife for interacting with EVM smart contracts, sending transactions and getting chain data.
- **Anvil**: Local Ethereum node, akin to Ganache, Hardhat Network.
- **Chisel**: Fast, utilitarian, and verbose solidity REPL.

## Documentation

https://book.getfoundry.sh/

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy

```shell
$ forge script script/Counter.s.sol:CounterScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Cast

```shell
$ cast <subcommand>
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```

# 安装openzeppelin-contracts

forge install OpenZeppelin/openzeppelin-contracts

# 安装特定版本

forge install OpenZeppelin/openzeppelin-contracts@v4.9.5

# 部署合约

source .env
forge script script/BeggingContract.s.sol:DeployBeggingContract --rpc-url sepolia --broadcast -vvv --private-key private_key

# 根据sol生成go语言代码（go语言交互）

sh script/generate_contract_go.sh BeggingContract

# 调用合约（cast交互）

## 调用donate方法

cast send 0xa65e5b57a3a93dfa904bd42c5b4d1e0d34ac3868 "donate()" --value $(cast --to-wei 0.1) --rpc-url https://sepolia.infura.io/v3/cb72dcaafb7c4728b720681d5345dfe4 --private-key xxxxx

## 查询总共捐赠了多少

cast call 0xa65e5b57a3a93dfa904bd42c5b4d1e0d34ac3868 "totalDonations()" --rpc-url https://sepolia.infura.io/v3/cb72dcaafb7c4728b720681d5345dfe4

## 查询某个地址捐赠了多少

cast call 0xa65e5b57a3a93dfa904bd42c5b4d1e0d34ac3868 "getDonation(address)(uint256)" 0xb052360268d7F2FA4Eb62eD1c6194257935C0BfE --rpc-url https://sepolia.infura.io/v3/cb72dcaafb7c4728b720681d5345dfe4