// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract MergeSortedArray {
    constructor(){

    }

    function merge(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {
        if (arr1.length == 0) {
            return arr2;
        }
        if (arr2.length == 0) {
            return arr1;
        }
        uint[] memory mergedArr = new uint[](arr1.length + arr2.length);
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                mergedArr[k++] = arr1[i++];
            } else {
                mergedArr[k++] = arr2[j++];
            }
        }
        while (i < arr1.length) {
            mergedArr[k++] = arr1[i++];
        }
        while (j < arr2.length) {
            mergedArr[k++] = arr2[j++];
        }
        return mergedArr;
    }

}
