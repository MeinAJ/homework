// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

// ✅ 创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数
contract Voting {
    mapping(address => uint64) public votes;
    mapping(address => bool) public voteStatus;
    address[] public candidates;
    address public owner;

    constructor(address _owner) {
        owner = _owner;
    }

    function vote(address candidate) public {
        require(!voteStatus[msg.sender], "Already voted");
        voteStatus[msg.sender] = true;
        // 仅在候选人首次被投票时加入数组
        if (votes[candidate] == 0) {
            candidates.push(candidate);
        }
        votes[candidate] += 1;
    }

    function getVotes(address candidate) public view returns (uint64) {
        //直接返回 mapping 值（默认返回0）
        return votes[candidate];
    }

    function resetVotes() public {
        _checkOnlyOwner();
        // 仅重置数组中的候选人票数
        for (uint i = 0; i < candidates.length;) {
            votes[candidates[i]] = 0;
            unchecked {++i;} // 禁用溢出检查节省 gas
        }
        // 清空数组但保留分配的内存 (更省 gas)
        delete candidates;
    }

    function _checkOnlyOwner() internal view {
        require(msg.sender == owner, "Only owner");
    }
}