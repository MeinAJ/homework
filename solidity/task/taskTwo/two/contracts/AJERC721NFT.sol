// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {ERC721URIStorage} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";


contract AJERC721NFT is ERC721URIStorage, Ownable {
    uint256 private _tokenIdCounter;

    constructor() Ownable(msg.sender) ERC721("AJERC721NFT", "AJERC721NFT") {

    }

    function mint(address to, string memory tokenURI) public onlyOwner {
        uint256 tokenId = _tokenIdCounter++;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
    }
}
