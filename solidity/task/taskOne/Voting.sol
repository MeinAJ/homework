// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

// ✅ 创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数
contract VotingContract {
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

// ✅ 反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
contract ReverseStringContract {
    function reverseString(string memory str) public pure returns (string memory) {
        // 1. 先将字符串转换为 bytes 数组
        // 2. 遍历数组，将每个字节反转
        // 3. 将反转后的字节数组转换为字符串
        // 这种只是针对ascii字符的反转，对于中文等其他字符的处理需要更复杂的算法
        bytes memory originalBytes = bytes(str);
        bytes memory reversedBytes = new bytes(originalBytes.length);
        for (uint i = 0; i < originalBytes.length; i++) {
            reversedBytes[i] = originalBytes[originalBytes.length - 1 - i];
        }
        return string(reversedBytes);
    }
}

// ✅  用 solidity 实现整数转罗马数字
//题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
contract RomanNumberConvert2IntContract {
    function RomanNumberConvert2Int(string memory str) public pure returns (uint64) {
        if (str == "") {
            return 0;
        }
        uint64 result = 0;
        // 先转成bytes数组
        bytes memory originalBytes = bytes(str);

        return 0;
    }
}
