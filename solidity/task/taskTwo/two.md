# ERC721 NFT

```
# 创建一个hardhat项目
npx hardhat init

# 安装openzeppelin依赖
npm install @openzeppelin/contracts

# 创建.env文件
touch .env

# 配置.env文件
INFURA_PROJECT_ID=xxx
INFURA_API_KEY_SECRET=xxx
WALLET_PRIVATE_KEY=xxx

# 安装dotenv依赖
npm install dotenv --save

# 配置hardhat.config.js
require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
    solidity: "0.8.28",
    networks: {
        sepolia: {
            url: `https://sepolia.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
            accounts: [process.env.WALLET_PRIVATE_KEY] // 使用钱包私钥
        }
    }
};

# 在contracts目录下创建AJERC721NFT.sol文件

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {ERC721URIStorage} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";


contract AjFirstNftImage is ERC721URIStorage, Ownable {
    uint256 private _tokenIdCounter;

    constructor() Ownable(msg.sender) ERC721("AJERC721NFT", "AJERC721NFT") {

    }

    function mint(address to, string memory tokenURI) public onlyOwner {
        uint256 tokenId = _tokenIdCounter++;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
    }
}


# 执行sol编译命令
npx hardhat compile

# 部署合约到sepolia网络
mkdir scripts
touch scripts/deploy.js

deploy.js内容如下：

const {ethers} = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contract with account:", deployer.address);

    const AjFirstNftImage = await ethers.getContractFactory("AJERC721NFT");
    const ajFirstNftImage = await AjFirstNftImage.deploy();

    await ajFirstNftImage.waitForDeployment();
    console.log("Contract deployed to:", await ajFirstNftImage.getAddress());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});

# 执行部署脚本
npx hardhat run scripts/deploy.js --network sepolia

# 安装ipfs-http-client
npm install ipfs-http-client

# 编写以下js脚本
# getMetadataUrl.js（获取NFT的元数据）
# uploadImage.js（上传文件到IPFS）
# uploadMetadata.js（上传NFT的元数据到IPFS）
# mintNFT.js（铸造NFT）

# 执行流程：
先执行getMetadataUrl.js，拿到metadataurl（公网的）：npx hardhat run scripts/getMetadataUrl.js

将metadataurl配置在mintNFT.js中，执行脚本：npx hardhat run scripts/mintNFT.js

```
