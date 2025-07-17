# ERC721 NFT

```
# 创建一个hardhat项目
npx hardhat init

# 安装dotenv依赖
nvm install 18.17.1  # 安装指定版本
nvm use 18.17.1      # 切换版本
nvm install --lts
nvm use --lts

npm install dotenv --save
npm install @nomicfoundation/hardhat-toolbox @openzeppelin/contracts @chainlink/contracts hardhat-deploy --verbose

# 创建.env文件
touch .env

# 配置.env文件
INFURA_PROJECT_ID=xxx
INFURA_API_KEY_SECRET=xxx
WALLET_PRIVATE_KEY=xxx

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

# 在contracts目录下创建AJAuctionNFT721.sol文件
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract AJAuctionNFT721 is ERC721URIStorage, Ownable {

    uint256 private _nextTokenId;

    constructor() ERC721("AJAuctionNFT721", "AJNFT721") Ownable(msg.sender) {

    }

    function mint(address to, string memory tokenURI) public onlyOwner {
        uint256 tokenId = _nextTokenId++;
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
    // 1. 获取部署者账户
    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);

    // 2. 获取合约工厂（自动读取 artifacts 中的 ABI 和字节码）
    const ContractFactory = await ethers.getContractFactory("AJAuctionNFT721");

    // 3. 部署合约（传入构造函数参数）
    const contract = await ContractFactory.deploy();

    // 4. 等待合约部署完成
    await contract.waitForDeployment();

    // 5. 获取合约地址
    const contractAddress = await contract.getAddress();
    console.log("合约地址:", contractAddress);

    return contractAddress;
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

# 执行部署脚本
npx hardhat run scripts/deploy.js --network sepolia

# 在pinata上上传图片获取到ipfs地址
1、https://app.pinata.cloud/ipfs/files，上传图片
2、点击图片，浏览器会自动跳转到ipfs地址，拿到地址如下：
https://moccasin-obliged-damselfly-991.mypinata.cloud/ipfs/bafybeicz7sn5y3yorsnajffdcuuspnkls6v4dc5cnbjxvfj5bxc45edcpi

# 创建metadata.json文件, 并上传到pinata上
1、image=https://moccasin-obliged-damselfly-991.mypinata.cloud/ipfs/bafybeicz7sn5y3yorsnajffdcuuspnkls6v4dc5cnbjxvfj5bxc45edcpi
metadata.json内容如下：
{
    "name": "AJAuctionNFT721",
    "description": "AJ's Auction NFT of ERC 721",
    "image": "https://moccasin-obliged-damselfly-991.mypinata.cloud/ipfs/bafybeicz7sn5y3yorsnajffdcuuspnkls6v4dc5cnbjxvfj5bxc45edcpi",
    "attributes": [
        {
            "trait_type": "Rarity",
            "value": "Legendary"
        },
        {
            "trait_type": "Color",
            "value": "Gold"
        }
    ]
}
2、得到metadata的ipfs地址，拿到地址如下：
https://moccasin-obliged-damselfly-991.mypinata.cloud/ipfs/bafkreid7fm6yctar4azcgvmvbab3kgz5trzp6ey2d4ujchf3wo5c47vlci

# 调用mint.js,创建NFT
1、touch scripts/mint.js
2、mint.js内容如下：
const {ethers} = require("hardhat");

async function mintNFT() {
    // 这里metadataUrl是一个公网可访问的链接
    const metadataUrl = "https://moccasin-obliged-damselfly-991.mypinata.cloud/ipfs/bafkreid7fm6yctar4azcgvmvbab3kgz5trzp6ey2d4ujchf3wo5c47vlci";
    const [owner] = await ethers.getSigners();

    const contractAddress = "0xcEeFAee09aA9C12827c5cD0654c4610C2A2b6492";
    const AjFirstNftImage = await ethers.getContractFactory("AJAuctionNFT721");
    const ajFirstNftImage = AjFirstNftImage.attach(contractAddress);

    // 铸造 NFT
    const tx = await ajFirstNftImage.mint(owner.address, metadataUrl);
    await tx.wait();
    console.log("NFT minted!");
}

mintNFT();

# 执行mint.js脚本，创建NFT
npx hardhat run scripts/mint.js --network sepolia
```
