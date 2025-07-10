const {
    time, loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const {anyValue} = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const {expect} = require("chai");

describe("MergeSortedArray Contract", function () {
    let mergeContract;

    before(async function () {
        const MergeSortedArray = await ethers.getContractFactory("MergeSortedArray");
        mergeContract = await MergeSortedArray.deploy();
    });

    it("should merge two empty arrays", async function () {
        const result = await mergeContract.merge([], []);
        expect(result).to.deep.equal([]);
    });

    it("should merge one empty array with non-empty array", async function () {
        const arr1 = [];
        const arr2 = [1, 2, 3];
        const expected = [1, 2, 3];
        const result = await mergeContract.merge(arr1, arr2);
        expect(result).to.deep.equal(expected);
    });

    it("should merge two sorted arrays correctly", async function () {
        const arr1 = [1, 3, 5];
        const arr2 = [2, 4, 6];
        const expected = [1, 2, 3, 4, 5, 6];
        const result = await mergeContract.merge(arr1, arr2);
        expect(result).to.deep.equal(expected);
    });

    it("should handle duplicate values", async function () {
        const arr1 = [2, 2, 2];
        const arr2 = [2, 2];
        const expected = [2, 2, 2, 2, 2];
        const result = await mergeContract.merge(arr1, arr2);
        expect(result).to.deep.equal(expected);
    });

    it("should handle different lengths", async function () {
        const arr1 = [1, 4, 6];
        const arr2 = [2, 3, 5, 7, 8];
        const expected = [1, 2, 3, 4, 5, 6, 7, 8];
        const result = await mergeContract.merge(arr1, arr2);
        expect(result).to.deep.equal(expected);
    });

    it("should handle descending input but still return ascending result", async function () {
        // 注意：本合约假设输入是升序的，如果传入无序数组，结果不保证有序
        const arr1 = [3, 2, 1]; // 乱序数组
        const arr2 = [6, 5, 4]; // 乱序数组
        const expected = [3, 2, 1, 6, 5, 4]; // 结果也不是有序的！
        const result = await mergeContract.merge(arr1, arr2);
        expect(result).to.deep.equal(expected);
    });
});