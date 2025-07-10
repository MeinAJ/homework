// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

// ✅ 创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数
contract Voting {

    //一个mapping来存储候选人的得票数
    mapping(address => uint64) public votes;
    mapping(address => bool) public voteStatus;
    address[] public candidates;
    address public owner;

    constructor(address _owner){
        owner = _owner;
    }

    //一个vote函数，允许用户投票给某个候选人，限制每个用户只能投一次
    function vote(address candidate) public {
        if (voteStatus[msg.sender] == false) {
            voteStatus[msg.sender] = true;
            votes[candidate] += 1;
            if (votes[candidate] > 0) {
                candidates.push(candidate);
            }
        }
    }

    //一个getVotes函数，返回某个候选人的得票数
    function getVotes(address candidate) public view returns (int64) {
        if (votes[candidate] > 0) {
            return votes[candidate];
        } else {
            return 0;
        }
    }

    //一个resetVotes函数，重置所有候选人的得票数
    //只能由owner调用
    function resetVotes() public {
        _checkOnlyOwner();
        for (uint i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
        delete candidates;
    }

    function _checkOnlyOwner() internal view {
        require(msg.sender == owner, "Only owner can call 'resetVotes()' function");
    }

}
