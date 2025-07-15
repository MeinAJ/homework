# ERC721 NFT

```
# 创建一个hardhat项目
npx hardhat init

# 安装dotenv依赖
npm install dotenv --save

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

# 在contracts目录下创建AJBeggingContract.sol文件

pragma solidity ^0.8.28;

contract AJBeggingContract {

    struct DonationInfo {
        address from;
        uint256 amount;
    }

    event DonationEvent(address from, uint256 amount, uint256 timestamp);

    uint256 private beginTimestamp;
    uint256 private endTimestamp;
    address private _owner;
    address[] private donators;
    mapping(address => uint256) private donationMapping;
    mapping(address => DonationInfo[]) private donationInfos;

    constructor(){
        _owner = msg.sender;
        beginTimestamp = block.timestamp;
        endTimestamp = block.timestamp + 365 * 1 days;
    }

    receive() public payable {
        _donate(msg.sender, msg.value);
        emit DonationEvent(msg.sender, msg.value, block.timestamp);
    }

    function _donate(address from, uint256 amount) internal {
        _checkTimestamp();
        donationMapping[from] += amount;
        donationInfos[from].push(DonationInfo(from, amount));
        if (donationMapping[from] > 0) {
            donators.push(from);
        }
    }

    function withdraw() public {
        _checkOwner();
        payable(msg.sender).transfer(address(this).balance);
    }

    function getDonation(address from) public view returns (uint256) {
        return donationMapping[from];
    }

    function findTop3Address(uint topCount) public view returns (address[]) {
        // 根据donators和donationMapping，先排序，再取前几个
        address[] memory result = new address[](topCount);
        uint256[] memory amounts = new uint256[](topCount);
        for (uint i = 0; i < donators.length; i++) {
            uint index = 0;
            for (uint j = 0; j < topCount; j++) {
                if (amounts[j] < donationMapping[donators[i]]) {
                    index = j;
                    break;
                }
            }
            if (index < topCount) {
                result[index] = donators[i];
                amounts[index] = donationMapping[donators[i]];
            }
        }
        return result;
    }

    function _checkTimestamp() internal pure {
        require(block.timestamp > beginTimestamp, "not started!");
        require(block.timestamp < endTimestamp, "already finished!");
    }

    function _checkOwner() internal pure {
        require(msg.sender == _owner, "need owner!");
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
    const ContractFactory = await ethers.getContractFactory("AJBeggingContract");

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

```
