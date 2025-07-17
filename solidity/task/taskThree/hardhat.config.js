require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
require('hardhat-deploy');
require("@openzeppelin/hardhat-upgrades")

module.exports = {
  solidity: "0.8.28",
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts: [process.env.WALLET_PRIVATE_KEY] // 使用钱包私钥
    },
    namedAccounts: {
      deployer: 0,
      user1: 1,
      user2: 2,
    }
  }
};