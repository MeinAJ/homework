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