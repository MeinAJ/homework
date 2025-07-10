// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {console} from "hardhat/console.sol";

contract Voting {
    constructor(){

    }
    // 创建一个mapping，key为候选人地址，value为投票数量
    mapping(address => uint) public candidateMapping;
    // 创建一个数组，用于存储候选人地址
    address[] public candidates;

    function vote(address candidate) public {
        // 判断candidate是否存在
        if (candidateMapping[candidate] == 0) {
            // 如果candidate不存在，则创建候选人
            candidateMapping[candidate] = 1;
            candidates.push(candidate);
        } else {
            // 如果candidate存在，则增加投票数量
            candidateMapping[candidate] += 1;
        }
    }

    function getVotes(address candidate) public view returns (uint){
        // 返回candidate的投票数量
        uint count = candidateMapping[candidate];
        return count;
    }

    function resetVotes() public {
        // 重置所有候选人的投票数量
        for (uint i = 0; i < candidates.length; i++) {
            candidateMapping[candidates[i]] = 0;
        }
    }
}
