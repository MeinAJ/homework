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
        for (uint i = 0; i < originalBytes.length;) {
            reversedBytes[i] = originalBytes[originalBytes.length - 1 - i];
            unchecked {i++;}
        }
        return string(reversedBytes);
    }
}

// ✅  用 solidity 实现整数转罗马数字
//题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
contract IntegerToRoman {
    // 定义罗马数字符号和对应的值（按从大到小排序）
    struct RomanNumeral {
        uint256 value;
        string symbol;
    }
    
    RomanNumeral[] private romanNumerals;

    constructor() {
        // 初始化罗马数字符号表（必须按值从大到小排序）
        romanNumerals.push(RomanNumeral(1000, "M"));
        romanNumerals.push(RomanNumeral(900, "CM"));
        romanNumerals.push(RomanNumeral(500, "D"));
        romanNumerals.push(RomanNumeral(400, "CD"));
        romanNumerals.push(RomanNumeral(100, "C"));
        romanNumerals.push(RomanNumeral(90, "XC"));
        romanNumerals.push(RomanNumeral(50, "L"));
        romanNumerals.push(RomanNumeral(40, "XL"));
        romanNumerals.push(RomanNumeral(10, "X"));
        romanNumerals.push(RomanNumeral(9, "IX"));
        romanNumerals.push(RomanNumeral(5, "V"));
        romanNumerals.push(RomanNumeral(4, "IV"));
        romanNumerals.push(RomanNumeral(1, "I"));
    }

    function intToRoman(uint256 num) public pure returns (string memory) {
        // 参数检查
        require(num > 0 && num < 4000, "Number must be between 1 and 3999");
        
        // 预定义罗马数字符号表（按值从大到小排序）
        RomanNumeral[13] memory numerals = [
            RomanNumeral(1000, "M"),
            RomanNumeral(900, "CM"),
            RomanNumeral(500, "D"),
            RomanNumeral(400, "CD"),
            RomanNumeral(100, "C"),
            RomanNumeral(90, "XC"),
            RomanNumeral(50, "L"),
            RomanNumeral(40, "XL"),
            RomanNumeral(10, "X"),
            RomanNumeral(9, "IX"),
            RomanNumeral(5, "V"),
            RomanNumeral(4, "IV"),
            RomanNumeral(1, "I")
        ];
        
        // 构建罗马数字字符串
        bytes memory roman;
        for (uint256 i = 0; i < numerals.length; i++) {
            while (num >= numerals[i].value) {
                roman = abi.encodePacked(roman, numerals[i].symbol);
                num -= numerals[i].value;
            }
        }
        
        return string(roman);
    }
}


// ✅  用 solidity 实现罗马数字转数整数
//题目描述在 https://leetcode.cn/problems/integer-to-roman/description/

contract RomanNumberConvert2IntContract {
    error InvalidRomanCharacter();
    function RomanNumberConvert2Int(string memory str) public pure returns (uint64) {
        bytes memory roman = bytes(str);
        uint256 len = roman.length;
        uint64 result = 0;
        uint64 current;
        uint64 next;
        for (uint256 i = 0; i < len; ) {
            current = _charToUint(roman[i]);
            // 检查是否有下一个字符
            if (i + 1 < len) {
                next = _charToUint(roman[i + 1]);
                // 组合数字规则：小值在左表示减法
                if (current < next) {
                    result += (next - current);
                    unchecked { i += 2; } // 跳过两个字符
                    continue;
                }
            }
            result += current;
            unchecked { i += 1; } // 处理单个字符
        }
        return result;
    }

    // 内部函数：罗马字符转数字
    function _charToUint(bytes1 char) internal pure returns (uint64) {
        if (char == 'I') return 1;
        if (char == 'V') return 5;
        if (char == 'X') return 10;
        if (char == 'L') return 50;
        if (char == 'C') return 100;
        if (char == 'D') return 500;
        if (char == 'M') return 1000;
        revert InvalidRomanCharacter();
    }
}

// ✅ 合并两个有序数组 (Merge Sorted Array) //题目描述：将两个有序数组合并为一个有序数组。
contract MergeSortedArrayContract {
    function MergeSortedArray(uint[] memory sortedArray1, uint[] memory sortedArray2) public pure returns (uint[] memory){
        // 思路
        // 1、sortedArray1和sortedArray2两两比较，小的元素直接放入result中，然后对应的数组index+1，直到两个数组中的一个数组遍历完。
        // 2、最后可能有一个数组没遍历完。
        uint[] memory result = new uint[](sortedArray1.length + sortedArray2.length);
        uint i = 0;
        uint j = 0;
        while (i < sortedArray1.length && j < sortedArray2.length) {
            if (sortedArray1[i] < sortedArray2[j]) {
                result.push(sortedArray1[i++]);
            } else {
                result.push(sortedArray2[j++]);
            }
        }
        while (i < sortedArray1.length) {
            result.push(sortedArray1[i++]);
        }
        while (j < sortedArray2.length) {
            result.push(sortedArray2[j++]);
        }
        return result;
    }
}
