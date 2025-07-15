// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract AJBeggingContract {

    struct DonationInfo {
        address from;
        uint256 amount;
    }

    event DonationEvent(address from, uint256 amount, uint256 timestamp);

    uint256 private beginTimestamp;
    uint256 private endTimestamp;
    address private _owner;
    address[] private donators;
    mapping(address => uint256) private donationMapping;
    mapping(address => DonationInfo[]) private donationInfos;

    constructor(){
        _owner = msg.sender;
        beginTimestamp = block.timestamp;
        endTimestamp = block.timestamp + 365 * 1 days;
    }

    receive() external payable {
        _donate(msg.sender, msg.value);
        emit DonationEvent(msg.sender, msg.value, block.timestamp);
    }

    function _donate(address from, uint256 amount) internal {
        _checkTimestamp();
        donationMapping[from] += amount;
        donationInfos[from].push(DonationInfo(from, amount));
        if (donationMapping[from] > 0) {
            donators.push(from);
        }
    }

    function withdraw() public {
        _checkOwner();
        payable(msg.sender).transfer(address(this).balance);
    }

    function getDonation(address from) public view returns (uint256) {
        return donationMapping[from];
    }

    function findTop3Address(uint topCount) public view returns (address[] memory) {
        // 根据donators和donationMapping，先排序，再取前几个
        address[] memory result = new address[](topCount);
        uint256[] memory amounts = new uint256[](topCount);
        for (uint i = 0; i < donators.length; i++) {
            uint index = 0;
            for (uint j = 0; j < topCount; j++) {
                if (amounts[j] < donationMapping[donators[i]]) {
                    index = j;
                    break;
                }
            }
            if (index < topCount) {
                result[index] = donators[i];
                amounts[index] = donationMapping[donators[i]];
            }
        }
        return result;
    }

    function _checkTimestamp() internal view {
        require(block.timestamp > beginTimestamp, "not started!");
        require(block.timestamp < endTimestamp, "already finished!");
    }

    function _checkOwner() internal view {
        require(msg.sender == _owner, "need owner!");
    }

}
