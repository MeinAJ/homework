// 导入uploadImage.js包
const ipfsHttpClient = require('ipfs-http-client');
const fs = require('fs'); // 新增此行
const {create} = ipfsHttpClient;
require('dotenv').config();

// 配置 Infura IPFS 客户端
const auth = 'Basic ' + Buffer.from(
    process.env.INFURA_PROJECT_ID + ':' + process.env.INFURA_API_KEY_SECRET
).toString('base64');

const ipfs = create({
    host: 'ipfs.infura.io',
    port: 5001,
    protocol: 'https',
    headers: {authorization: auth},
});

async function uploadMetadata(imageUrl) {
    const metadata = {
        name: "AJERC721NFT",
        description: "AJERC721NFT",
        image: imageUrl,
        attributes: [
            {trait_type: "Rarity", value: "Legendary"}
        ]
    };

    try {
        // 上传元数据
        const {cid} = await ipfs.add(JSON.stringify(metadata));
        const metadataUrl = `ipfs://${cid}`;
        console.log('Metadata uploaded to IPFS. URL:', metadataUrl);
        return metadataUrl;
    } catch (error) {
        console.error('Error uploading metadata:', error);
    }
}

module.exports = {
    uploadMetadata
};