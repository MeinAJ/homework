# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a Hardhat Ignition module that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat service
REPORT_GAS=true npx hardhat service
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/Lock.js
```
# 创建一个hardhat项目
npx hardhat init

# 编译
npx hardhat compile

# 测试
npx hardhat test

# 运行部署脚本
npx hardhat run scripts/erc_deploy.js --network sepolia

npx hardhat run scripts/nft_deploy.js --network sepolia

npx hardhat run scripts/begging_deploy.js --network sepolia

# 锻造
npx hardhat run scripts/mintNFT.js --network sepolia

