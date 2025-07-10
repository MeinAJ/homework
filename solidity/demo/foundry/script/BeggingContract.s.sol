pragma solidity ^0.8.28;

import {CommonBase} from "../lib/forge-std/src/Base.sol";
import {Script} from "../lib/forge-std/src/Script.sol";
import {StdChains} from "../lib/forge-std/src/StdChains.sol";
import {StdCheatsSafe} from "../lib/forge-std/src/StdCheats.sol";
import {StdUtils} from "../lib/forge-std/src/StdUtils.sol";
import {BeggingContract} from "../src/BeggingContract.sol";

contract DeployBeggingContract is Script {

    BeggingContract public beggingContract;

    function run() external {
        vm.startBroadcast();
        beggingContract = new BeggingContract();
        vm.stopBroadcast();
    }
}