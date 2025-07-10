// scripts/uploadMetadata.js
const {uploadMetadata} = require("./uploadMetadata");
const {uploadImage} = require("./uploadImage");

async function main() {
    // 1. 上传图片到 IPFS
    const imageUrl = await uploadImage("/Users/ja/Desktop/ipfs.png");
    // 2. 生成并上传元数据
    const metadataUrl = await uploadMetadata(imageUrl);
    console.log("Metadata URL:", metadataUrl);
}

main().catch((error) => {
    console.error(error);
    process.exit(1);
});