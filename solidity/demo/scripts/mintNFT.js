// 铸造 NFT
const {ethers} = require("hardhat");

async function mintNFT() {
    // 这里metadataUrl是一个公网可访问的链接
    const metadataUrl = "https://moccasin-obliged-damselfly-991.mypinata.cloud/ipfs/bafkreihkiekitcdsbkqbwc5a3ian7bli6nt44s26snfzrv7zearwd2tsra";
    const [owner] = await ethers.getSigners();

    const contractAddress = "0x89780d8b832E11c98792Fa8a5ffCDe74311497F4";
    const AjFirstNftImage = await ethers.getContractFactory("AjFirstNftImage");
    const ajFirstNftImage = AjFirstNftImage.attach(contractAddress);

    // 铸造 NFT
    const tx = await ajFirstNftImage.mint(owner.address, metadataUrl);
    await tx.wait();
    console.log("NFT minted!");
}

mintNFT();