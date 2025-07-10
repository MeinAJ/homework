// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

// 捐赠合约
contract BeggingContract is Ownable {

    // 捐赠事件
    event DonationReceived(address indexed donor, uint amount, uint timestamp);

    // 捐赠信息struct
    struct Donation {
        address donor;
        uint amount;
        uint timestamp;
    }

    // 捐赠期限, 10年
    uint private beginDonationTime;

    // 捐赠结束时间
    uint private endDonationTime;

    // 捐赠者地址 => 捐赠金额
    mapping(address => uint) public donationMapping;

    // 捐赠信息列表
    Donation[] public donations;

    // 构造函数
    constructor() Ownable(msg.sender) {
        beginDonationTime = block.timestamp;
        endDonationTime = block.timestamp + 10 * 365 days;
    }

    // 检查时间
    function _checkTime() internal view {
        require(block.timestamp >= beginDonationTime, "BeggingContract: donation not started yet");
        require(block.timestamp <= endDonationTime, "BeggingContract: donation ended");
    }

    // 检查捐赠金额
    function _checkAmount() internal view {
        require(msg.value >= 1 wei, "BeggingContract: donation amount should be at least 1 wei");
    }

    // 捐赠
    function donate() external payable {
        _checkTime();
        _checkAmount();
        // 记录捐赠信息
        Donation memory donation = Donation(msg.sender, msg.value, block.timestamp);
        donations.push(donation);
        donationMapping[msg.sender] += msg.value;
        // 触发捐赠事件
        emit DonationReceived(msg.sender, msg.value, block.timestamp);
    }

    // 提取所有捐赠
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        (bool success,) =  payable(owner()).call{value: balance}("");
        require(success, "Transfer failed");
    }

    // 获取某个捐赠者的捐赠金额
    function getDonation(address from) external view returns (uint) {
        return donationMapping[from];
    }

    // 获取所有捐赠信息
    function totalDonations() public view returns (uint) {
        return address(this).balance;
    }

}