// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract ReverseString {
    constructor(){

    }

    function reverseString(string memory str) public pure returns (string memory) {
        // 反转字符串
        bytes memory strBytes = bytes(str);
        bytes memory reversedBytes = new bytes(strBytes.length);
        for (uint i = 0; i < strBytes.length; i++) {
            reversedBytes[i] = strBytes[strBytes.length - 1 - i];
        }
        return string(reversedBytes);
    }

}
