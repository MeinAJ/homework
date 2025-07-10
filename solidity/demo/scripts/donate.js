const {ethers} = require("hardhat");

async function main() {
    const [signer] = await ethers.getSigners();

    const contractAddress = "0x3beec3FC5aD3771A8856927200e98D28473721B7"; // 替换为部署后的地址
    const amountToDonate = "0.1"; // 0.1 ETH

    const Donate = await ethers.getContractAt("BeggingContract", contractAddress, signer);

    console.log(`Sending donation of ${amountToDonate} ETH...`);
    const tx = await Donate.donate({
        value: ethers.parseEther(amountToDonate)
    });

    await tx.wait();
    console.log("Donation successful!");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });