// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20, Ownable {
    uint256 public constant RATE = 100000000; // 100000000 MyToken per 1 ETH
    uint256 public constant MIN_ETH = 0.001 ether;

    constructor() ERC20("RCCDemoToken", "RDT") Ownable(msg.sender) {
    }

    function mint() public payable {
        require(msg.value >= MIN_ETH, "Not enough ETH sent");
        uint256 tokensToMint = (msg.value * RATE);
        _mint(msg.sender, tokensToMint);
    }

    function withdrawETH() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No ETH to withdraw");
        payable(owner()).transfer(balance);
    }

    function getTokens(address _from) public view returns (uint256) {
        return balanceOf(_from);
    }

    receive() external payable {
        mint();
    }
}