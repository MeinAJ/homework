// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract AJAuctionNFT721 is ERC721URIStorage, Ownable {

    uint256 private _nextTokenId;

    constructor() ERC721("AJAuctionNFT721", "AJNFT721") Ownable(msg.sender) {

    }

    function mint(address to, string memory tokenURI) public onlyOwner {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
    }
}