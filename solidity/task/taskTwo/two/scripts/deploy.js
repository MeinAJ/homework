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