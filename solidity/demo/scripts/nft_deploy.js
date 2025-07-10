// scripts/deploy.js
const {ethers} = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contract with account:", deployer.address);

    const AjFirstNftImage = await ethers.getContractFactory("AjFirstNftImage");
    const ajFirstNftImage = await AjFirstNftImage.deploy();

    await ajFirstNftImage.waitForDeployment();
    console.log("Contract deployed to:", await ajFirstNftImage.getAddress());
    // address: 0x734880Ed49A1c3206DF342B88CbE78d10c85360F
    // address: 0xeD21a229933F67671190cC2424a713A7DfEF4BE2
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});