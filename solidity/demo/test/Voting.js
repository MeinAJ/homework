const {
    time,
    loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const {anyValue} = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const {expect} = require("chai");

describe("Voting Contract", function () {
    let Voting;
    let voting;
    let owner;
    let addr1;
    let addr2;

    beforeEach(async function () {
        // 获取合约工厂和账户
        Voting = await ethers.getContractFactory("Voting");
        [owner, addr1, addr2] = await ethers.getSigners();
        // 部署合约
        voting = await Voting.deploy();
    });

    describe("测试投票", function () {
        it("投票成功", async function () {
            // 给投票人投票
            await voting.vote(addr1.address);
            await voting.vote(addr2.address);
            // 打印投票结果
            console.log("addr1 voteCount", await voting.getVotes(addr1.address));
            console.log("addr2 voteCount", await voting.getVotes(addr2.address));
            // 重置投票
            await voting.resetVotes();
            // 打印投票结果
            console.log("addr1 voteCount", await voting.getVotes(addr1.address));
            console.log("addr2 voteCount", await voting.getVotes(addr2.address));
        });
    });

});