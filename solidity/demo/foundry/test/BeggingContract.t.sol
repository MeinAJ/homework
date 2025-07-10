// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {CommonBase} from "../lib/forge-std/src/Base.sol";
import {StdAssertions} from "../lib/forge-std/src/StdAssertions.sol";
import {StdChains} from "../lib/forge-std/src/StdChains.sol";
import {StdCheats, StdCheatsSafe} from "../lib/forge-std/src/StdCheats.sol";
import {StdUtils} from "../lib/forge-std/src/StdUtils.sol";
import {Test} from "../lib/forge-std/src/Test.sol";
import {console} from "../lib/forge-std/src/console.sol";
import {BeggingContract} from "../src/BeggingContract.sol";


contract BeggingContractTest is Test {
    BeggingContract public beggingContract;

    // 测试地址
    address owner = address(0xABCD);
    address donor1 = address(0x1);
    address donor2 = address(0x2);
    address attacker = address(0x999);

    // 测试金额
    uint256 donationAmount = 1 ether;
    uint256 smallDonation = 0.5 ether;

    // 时间常量
    uint256 constant TEN_YEARS = 10 * 365 days;
    uint256 startTime;

    function setUp() public {
        // 设置区块时间戳
        startTime = block.timestamp;
        vm.warp(startTime);

        // 使用特定地址部署合约
        vm.prank(owner);
        beggingContract = new BeggingContract();

        // 为捐赠者提供资金
        vm.deal(donor1, 10 ether);
        vm.deal(donor2, 10 ether);
        vm.deal(attacker, 10 ether);
    }

    // ======== 基础功能测试 ========

    // 测试合约初始化状态
    function test_InitialState() public view {
        assertEq(beggingContract.owner(), owner);
        assertEq(address(beggingContract).balance, 0);
        assertEq(beggingContract.totalDonations(), 0);
    }

    // 测试正常捐赠
    function test_Donate() public {

        // 验证事件
        vm.expectEmit(true, true, true, true);
        emit BeggingContract.DonationReceived(donor1, donationAmount, block.timestamp);

        // 捐赠
        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        // 验证捐赠记录
        assertEq(beggingContract.getDonation(donor1), donationAmount);
        assertEq(address(beggingContract).balance, donationAmount);
        assertEq(beggingContract.totalDonations(), donationAmount);

    }

    // 测试多次捐赠
    function test_MultipleDonations() public {
        vm.startPrank(donor1);
        beggingContract.donate{value: smallDonation}();
        beggingContract.donate{value: smallDonation}();
        vm.stopPrank();

        vm.prank(donor2);
        beggingContract.donate{value: donationAmount}();

        assertEq(beggingContract.getDonation(donor1), smallDonation * 2);
        assertEq(beggingContract.getDonation(donor2), donationAmount);
        assertEq(address(beggingContract).balance, smallDonation * 2 + donationAmount);
    }

    // ======== 边界条件测试 ========

    // 测试最小捐赠金额 (1 wei)
    function test_MinimumDonation() public {
        vm.prank(donor1);
        beggingContract.donate{value: 1 wei}();

        assertEq(beggingContract.getDonation(donor1), 1 wei);
    }

    // 测试捐赠金额不足
    function test_DonateWithZeroValue() public {
        vm.prank(donor1);
        vm.expectRevert("BeggingContract: donation amount should be at least 1 wei");
        beggingContract.donate{value: 0}();
    }

    // 测试捐赠开始前
    function test_DonateBeforeStart() public {
        // 回退到开始时间之前
        // 打印开始时间
        console.log("start time: %s", startTime);
        // 回退到开始时间之前
        vm.warp(0);

        vm.prank(donor1);
        vm.expectRevert("BeggingContract: donation not started yet");
        beggingContract.donate{value: donationAmount}();
    }

    // 测试捐赠结束后
    function test_DonateAfterEnd() public {
        // 快进到结束时间之后
        vm.warp(startTime + TEN_YEARS + 1 days);

        vm.prank(donor1);
        vm.expectRevert("BeggingContract: donation ended");
        beggingContract.donate{value: donationAmount}();
    }

    // ======== 提取功能测试 ========

    // 测试所有者提取资金
    function test_WithdrawByOwner() public {
        // 先捐赠
        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        uint256 contractBalance = address(beggingContract).balance;

        // 验证所有者可以提取
        vm.prank(owner);
        beggingContract.withdraw();

        assertEq(address(beggingContract).balance, 0);
        assertEq(address(owner).balance, contractBalance);
    }

    // 测试非所有者尝试提取
    function test_WithdrawByNonOwner() public {
        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", attacker));
        beggingContract.withdraw();

        assertEq(address(beggingContract).balance, donationAmount);
    }

    // 测试提取空合约
    function test_WithdrawEmptyContract() public {
        vm.prank(owner);
        vm.expectRevert("No funds to withdraw");
        beggingContract.withdraw();
    }

    // ======== 时间边界测试 ========

    // 测试捐赠开始时间边界
    function test_DonateAtStartTime() public {
        vm.warp(startTime);

        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        assertEq(beggingContract.getDonation(donor1), donationAmount);
    }

    // 测试捐赠结束时间边界
    function test_DonateAtEndTime() public {
        vm.warp(startTime + TEN_YEARS);

        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        assertEq(beggingContract.getDonation(donor1), donationAmount);
    }

    // ======== 捐赠列表测试 ========

    // 测试捐赠列表记录
    function test_DonationList() public {
        vm.prank(donor1);
        beggingContract.donate{value: smallDonation}();

        vm.prank(donor2);
        beggingContract.donate{value: donationAmount}();

        vm.prank(donor1);
        beggingContract.donate{value: smallDonation}();

        // 验证捐赠列表
        (address donor1Addr, uint256 amount1,) = beggingContract.donations(0);
        (address donor2Addr, uint256 amount2,) = beggingContract.donations(1);
        (address donor1Again, uint256 amount3,) = beggingContract.donations(2);

        assertEq(donor1Addr, donor1);
        assertEq(amount1, smallDonation);
        assertEq(donor2Addr, donor2);
        assertEq(amount2, donationAmount);
        assertEq(donor1Again, donor1);
        assertEq(amount3, smallDonation);
    }

    // 测试未捐赠地址查询
    function test_GetNonDonor() public view {
        assertEq(beggingContract.getDonation(attacker), 0);
    }

    // ======== 高级安全测试 ========

    // 测试合约接收ETH能力
    function test_ContractCannotReceiveETH() public {
        // 直接发送ETH到合约地址
        vm.prank(donor1);
        (bool success,) = address(beggingContract).call{value: donationAmount}("");
        assertTrue(success, "Should be able to receive ETH");

        // 但不会记录为捐赠
        assertEq(beggingContract.getDonation(donor1), 0);
        assertEq(address(beggingContract).balance, donationAmount);
    }

    // 测试捐赠后立即结束
    function test_DonateAndImmediatelyEnd() public {
        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        // 快进到结束时间之后
        vm.warp(startTime + TEN_YEARS + 1);

        // 尝试再次捐赠
        vm.prank(donor1);
        vm.expectRevert("BeggingContract: donation ended");
        beggingContract.donate{value: donationAmount}();

        // 提取资金
        uint256 contractBalance = address(beggingContract).balance;
        vm.prank(owner);
        beggingContract.withdraw();

        assertEq(address(beggingContract).balance, 0);
        assertEq(address(owner).balance, contractBalance);
    }

}