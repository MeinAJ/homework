// SPDX-License-Identifier: UNLICENSED
// 声明代码未授权任何许可

pragma solidity ^0.8.28;
// 指定Solidity编译器版本必须为0.8.28或更高(但不包括0.9.0及以上)

// 基于时间窗口的铸造控制合约
contract DynamicMinter {
    // 私有映射，记录哪些地址有铸造权限
    mapping(address => bool) private _minters;

    // 公共变量，每日铸造限额(初始1000 ether)
    uint256 public dailyMintLimit = 1000 ether;

    // 修饰器：限制只有有铸造权限的地址可以调用
    modifier onlyMinter() {
        // 检查调用者是否在_minters映射中标记为true
        require(_minters[msg.sender], "Minting denied");
        // 如果通过检查，继续执行函数
        _;
    }

    // 设置铸造权限状态函数(仅合约所有者可调用)
    function setMinter(address account, bool status) external onlyOwner {
        // 更新指定账户的铸造权限状态
        _minters[account] = status;
    }

    // 铸造函数(仅限有权限的地址调用)
    function mint(address to, uint256 amount) external onlyMinter {
        // 检查铸造数量不超过每日限额
        require(amount <= dailyMintLimit, "Exceed daily limit");
        // 调用内部铸造函数
        _mint(to, amount);
    }
}

// 结合链下签名验证的防女巫攻击铸造合约
contract SybilResistantMint {
    // 为bytes32类型添加ECDSA签名验证功能
    using ECDSA for bytes32;

    // 公共映射，记录每个地址的铸造次数
    mapping(address => uint256) public mintCounts;

    // 验证签名后铸造的函数
    function verifiedMint(
        uint256 amount,       // 要铸造的数量
        bytes memory signature, // 链下生成的签名
        address referrer      // 推荐人地址
    ) external {
        // 生成消息哈希(调用者地址+推荐人地址)
        bytes32 hash = keccak256(abi.encodePacked(msg.sender, referrer));

        // 验证签名是否由合约所有者生成
        require(
            hash.recover(signature) == owner(),
            "Invalid signature"
        );

        // 检查调用者铸造次数是否小于3次
        require(mintCounts[msg.sender] < 3, "Mint limit reached");

        // 增加调用者的铸造计数
        mintCounts[msg.sender]++;
        // 执行实际铸造
        _mint(msg.sender, amount);
    }
}

// 分级税率控制系统合约
contract TaxController {
    // 税率等级结构体定义
    struct TaxTier {
        uint256 minAmount;  // 适用该等级的最小金额
        uint256 feePercent; // 费用百分比
    }

    // 公共税率等级数组
    TaxTier[] public taxTiers;

    // 计算税费的内部视图函数
    function _calculateTax(uint256 amount) internal view returns (uint256) {
        // 从最高等级开始向下检查
        for (uint i = taxTiers.length; i > 0; i--) {
            // 如果金额大于等于该等级的最小金额
            if (amount >= taxTiers[i - 1].minAmount) {
                // 计算并返回税费(金额*百分比/100)
                return amount * taxTiers[i - 1].feePercent / 100;
            }
        }
        // 不满足任何等级则返回0税费
        return 0;
    }

    // 应用税费的内部函数
    function applyTax(address from, uint256 amount) internal {
        // 计算税费
        uint256 tax = _calculateTax(amount);
        // 如果税费大于0
        if (tax > 0) {
            // 将税费从发送者转到税费池
            _transfer(from, taxPool, tax);
        }
    }
}

// 可配置的资金分配合约
contract FundDistributor {
    // 资金分配结构体定义
    struct Allocation {
        address receiver;  // 接收地址
        uint16 percentage; // 分配百分比
    }

    // 公共分配方案数组
    Allocation[] public allocations;

    // 分配资金的内部函数
    function _distributeFunds(uint256 totalTax) internal {
        // 记录剩余资金(初始为全部税费)
        uint256 remaining = totalTax;

        // 遍历所有分配方案
        for (uint i = 0; i < allocations.length; i++) {
            // 计算当前接收者应得金额
            uint256 amount = totalTax * allocations[i].percentage / 100;
            // 执行转账
            _transfer(address(this), allocations[i].receiver, amount);
            // 扣除已分配金额
            remaining -= amount;
        }

        // 如果还有剩余资金
        if (remaining > 0) {
            // 将剩余资金转给默认接收者
            _transfer(address(this), defaultReceiver, remaining);
        }
    }
}

// Uniswap自动配对合约
contract AutoLiquidity {
    // 不可变的Uniswap V2路由器接口实例
    IUniswapV2Router02 public immutable uniswapRouter;

    // 添加流动性的内部函数
    function addLiquidity(uint256 tokenAmount, uint256 ethAmount) internal {
        // 批准路由器使用合约中的代币
        _approve(address(this), address(uniswapRouter), tokenAmount);

        // 调用Uniswap路由器的添加流动性方法
        uniswapRouter.addLiquidityETH{value: ethAmount}(
            address(this),    // 代币合约地址
            tokenAmount,     // 要添加的代币数量
            0,               // 代币数量最小接受值(滑点保护)
            0,               // ETH数量最小接受值(滑点保护)
            liquidityLock,   // 流动性锁定地址
            block.timestamp  // 交易截止时间
        );
    }

    // 锁定流动性的函数(仅所有者可调用)
    function lockLiquidity(uint256 daysToLock) external onlyOwner {
        // 检查锁定天数不超过365天
        require(daysToLock <= 365, "Lock period too long");
        // 计算锁定结束时间戳(当前时间+指定天数)
        liquidityLock = block.timestamp + daysToLock * 1 days;
    }
}

// 防止价格操控的DEX保护合约
contract DexProtection {
    // 修饰器：验证交易对是否在白名单中
    modifier validateDexPair(address pair) {
        require(isWhitelistedPair[pair], "Invalid DEX pair");
        _;
    }

    // 安全交换的内部函数
    function _safeSwap(
        address pair,           // 交易对地址
        uint256 amountOutMin,   // 最小输出量
        address[] memory path   // 交易路径
    ) internal validateDexPair(pair) {  // 使用修饰器验证交易对
        // 获取预期的交换输出量
        uint256[] memory amounts = uniswapRouter.getAmountsOut(msg.value, path);
        // 验证输出量满足最小要求
        require(amounts[1] >= amountOutMin, "Insufficient output");

        // 调用交易对的swap方法执行交换
        (bool success,) = pair.call{value: msg.value}(
        abi.encodeWithSignature(
        "swap(uint256,uint256,address,bytes)",
        amounts[0],     // 输入量
        amounts[1],    // 输出量
        msg.sender,     // 接收者
        new bytes(0)   // 额外数据
        );
        // 检查交换是否成功
        require(success, "Swap failed");
    }
}