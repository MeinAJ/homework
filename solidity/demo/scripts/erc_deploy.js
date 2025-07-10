const { ethers } = require("hardhat");

async function main() {
    // 1. 获取部署者账户
    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);

    // 2. 获取合约工厂（自动读取 artifacts 中的 ABI 和字节码）
    const ContractFactory = await ethers.getContractFactory("AjToken");

    // 3. 部署合约（传入构造函数参数，例如 value = 100）
    const contract = await ContractFactory.deploy(10000000);

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