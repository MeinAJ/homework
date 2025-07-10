// uploadImage.js
const ipfsHttpClient = require('ipfs-http-client');
const fs = require('fs'); // 新增此行
const {create} = ipfsHttpClient;
require('dotenv').config();

// 配置 Infura IPFS 客户端
const auth = 'Basic ' + Buffer.from(
    "cb72dcaafb7c4728b720681d5345dfe4" + ':' + "tT4g+BiH3L0KAOJ0C+tmqHrA2EKQJg9jDvZjo2cyFwdbUkOcK0toxQ"
).toString('base64');

const ipfs = create({
    host: 'ipfs.infura.io',
    port: 5001,
    protocol: 'https',
    headers: {authorization: auth},
});

async function uploadImage(filePath) {
    try {
        // 上传图片文件
        const file = fs.readFileSync(filePath);
        const {cid} = await ipfs.add(file);
        const imageUrl = `ipfs://${cid}`;
        console.log('Image uploaded to IPFS. CID:', cid.toString());
        return imageUrl;
    } catch (error) {
        console.error('Error uploading image:', error);
    }
}

module.exports = {
    uploadImage,
};