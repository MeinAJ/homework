const {
    time,
    loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const {anyValue} = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const {expect} = require("chai");

describe("ReverseString Contract", function () {
    let reverseString;

    // 在每个测试之前部署合约
    beforeEach(async function () {
        const ReverseString = await ethers.getContractFactory("ReverseString");
        reverseString = await ReverseString.deploy();
    });

    // 测试正常英文字符串
    it("should reverse a normal string", async function () {
        const input = "hello world";
        const expected = "dlrow olleh";
        const result = await reverseString.reverseString(input);
        expect(result).to.equal(expected);
    });

    // 测试单个字符
    it("should handle single character", async function () {
        const input = "a";
        const expected = "a";
        const result = await reverseString.reverseString(input);
        expect(result).to.equal(expected);
    });

    // 测试空字符串
    it("should handle empty string", async function () {
        const input = "";
        const expected = "";
        const result = await reverseString.reverseString(input);
        expect(result).to.equal(expected);
    });

    // 测试数字组成的字符串
    it("should reverse numeric string", async function () {
        const input = "1234567890";
        const expected = "0987654321";
        const result = await reverseString.reverseString(input);
        expect(result).to.equal(expected);
    });

    // 测试特殊符号
    it("should reverse special characters", async function () {
        const input = "!@#$%^&*()";
        const expected = ")(*&^%$#@!";
        const result = await reverseString.reverseString(input);
        expect(result).to.equal(expected);
    });

});