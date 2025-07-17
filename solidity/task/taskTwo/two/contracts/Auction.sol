// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/access/OwnableUpgradeable.sol";

contract Auction is UUPSUpgradeable, OwnableUpgradeable {
    IERC721 public nft;
    uint256 public tokenId;
    address public seller;
    uint256 public endTime;

    uint256 public highestBid;
    address public highestBidder;
    mapping(address => uint256) public pendingReturns;

    bool public ended;
    address public paymentToken;
    AggregatorV3Interface internal priceFeed;

    event NewHighestBid(address bidder, uint256 amount);
    event AuctionEnded(address winner, uint256 amount);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _nft,
        uint256 _tokenId,
        uint256 _duration,
        address _paymentToken,
        address _priceFeed
    ) external initializer {
        __Ownable_init(msg.sender);
        nft = IERC721(_nft);
        tokenId = _tokenId;
        seller = msg.sender;
        endTime = block.timestamp + _duration;
        paymentToken = _paymentToken;
        priceFeed = AggregatorV3Interface(_priceFeed);

        nft.transferFrom(seller, address(this), tokenId);
    }

    function bid(uint256 amount) external payable {
        require(block.timestamp < endTime, "Auction ended");
        require(amount > highestBid, "Bid too low");

        if (paymentToken == address(0)) {
            require(msg.value == amount, "ETH amount mismatch");
        } else {
            IERC20(paymentToken).transferFrom(msg.sender, address(this), amount);
        }

        // 使用 Chainlink 转换价格
        (, int price,,,) = priceFeed.latestRoundData();
        uint256 usdValue = (amount * uint256(price)) / 1e18;

        // 动态手续费逻辑
        uint256 fee = calculateFee(usdValue);
        uint256 netAmount = amount - fee;

        if (highestBidder != address(0)) {
            pendingReturns[highestBidder] += highestBid;
        }

        highestBid = netAmount;
        highestBidder = msg.sender;

        emit NewHighestBid(msg.sender, netAmount);
    }

    function endAuction() external {
        require(block.timestamp >= endTime, "Auction not ended");
        require(!ended, "Auction already ended");

        ended = true;

        if (highestBidder != address(0)) {
            nft.safeTransferFrom(address(this), highestBidder, tokenId);

            // 支付给卖家
            if (paymentToken == address(0)) {
                payable(seller).transfer(highestBid);
            } else {
                IERC20(paymentToken).transfer(seller, highestBid);
            }
        } else {
            nft.safeTransferFrom(address(this), seller, tokenId);
        }

        emit AuctionEnded(highestBidder, highestBid);
    }

    // 动态手续费计算
    function calculateFee(uint256 amount) internal pure returns (uint256) {
        if (amount < 1000e18) return amount * 5 / 100; // 5%
        if (amount < 10000e18) return amount * 3 / 100; // 3%
        return amount * 1 / 100; // 1%
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}