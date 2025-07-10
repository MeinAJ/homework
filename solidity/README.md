# solidity基础
# solidity进阶
# openZeppelin合约库（erc20、erc721）
# hardhat工具链
# foundry工具链
# 合约升级机制&预言机接入


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
