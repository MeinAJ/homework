// 导入uploadImage.js包
async function uploadMetadata(imageUrl) {
    const metadata = {
        name: "My First Image of NFT",
        description: "An NFT created with Infura IPFS",
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