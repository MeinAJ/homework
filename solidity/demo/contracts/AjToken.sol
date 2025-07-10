// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract AjToken is ERC20 {

    constructor(uint256 initialSupply) ERC20("AjToken", "AJT") {
        _mint(msg.sender, initialSupply);
    }

}
