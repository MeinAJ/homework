package taskTwo

// 作业 1：ERC20 代币
//任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
//合约包含以下标准 ERC20 功能：
//balanceOf：查询账户余额。
//transfer：转账。
//approve 和 transferFrom：授权和代扣转账。
//使用 event 记录转账和授权操作。
//提供 mint 函数，允许合约所有者增发代币。
//提示：
//使用 mapping 存储账户余额和授权信息。
//使用 event 定义 Transfer 和 Approval 事件。
//部署到sepolia 测试网，导入到自己的钱包

pragma solidity ^0.8.28;

interface IERC20 {
    // 核心函数
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address from, address to, uint256 amount) external returns (bool);

    // 事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract ERC20 is IERC20 {
    // 状态变量
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;
    string private _name;
    string private _symbol;

    // 构造函数（初始化代币名称、符号和总供应量）
    constructor(string memory name_, string memory symbol_, uint256 initialSupply_) {
        _name = name_;
        _symbol = symbol_;
        _mint(msg.sender, initialSupply_); // 初始代币铸造给部署者
    }

    // 1. 查询代币名称
    function name() public view returns (string memory) {
        return _name;
    }

    // 2. 查询代币符号（如 "ETH"）
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    // 3. 查询代币小数位数（默认18，兼容大多数钱包）
    function decimals() public pure returns (uint8) {
        return 18;
    }

    // 4. 查询总供应量（IERC20 必须实现）
    function totalSupply() public view override returns (uint256) {
        return _totalSupply;
    }

    // 5. 查询地址余额（IERC20 必须实现）
    function balanceOf(address account) public view override returns (uint256) {
        return _balances[account];
    }

    // 6. 转账（IERC20 必须实现）
    function transfer(address to, uint256 amount) public override returns (bool) {
        _transfer(msg.sender, to, amount);
        return true;
    }

    // 7. 查询授权额度（IERC20 必须实现）
    function allowance(address owner, address spender) public view override returns (uint256) {
        return _allowances[owner][spender];
    }

    // 8. 授权其他地址使用代币（IERC20 必须实现）
    function approve(address spender, uint256 amount) public override returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    // 9. 授权转账（IERC20 必须实现）
    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        _spendAllowance(from, msg.sender, amount); // 检查并扣除授权额度
        _transfer(from, to, amount);
        return true;
    }

    // === 内部函数 ===
    // 代币转账逻辑
    function _transfer(address from, address to, uint256 amount) internal {
        require(from != address(0), "ERC20: transfer from zero address");
        require(to != address(0), "ERC20: transfer to zero address");

        uint256 fromBalance = _balances[from];
        require(fromBalance >= amount, "ERC20: insufficient balance");
        unchecked {
            _balances[from] = fromBalance - amount;
            _balances[to] += amount;
        }

        emit Transfer(from, to, amount); // 触发转账事件
    }

    // 代币铸造（仅合约内部可调用）
    function _mint(address account, uint256 amount) internal {
        require(account != address(0), "ERC20: mint to zero address");
        _totalSupply += amount;
        _balances[account] += amount;
        emit Transfer(address(0), account, amount); // 从零地址铸造
    }

    // 授权逻辑
    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "ERC20: approve from zero address");
        require(spender != address(0), "ERC20: approve to zero address");
        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount); // 触发授权事件
    }

    // 检查并扣除授权额度
    function _spendAllowance(address owner, address spender, uint256 amount) internal {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= amount, "ERC20: insufficient allowance");
            unchecked {
                _approve(owner, spender, currentAllowance - amount);
            }
        }
    }
}

// 作业2：在测试网上发行一个图文并茂的 NFT
//任务目标
//使用 Solidity 编写一个符合 ERC721 标准的 NFT 合约。
//将图文数据上传到 IPFS，生成元数据链接。
//将合约部署到以太坊测试网（如 Goerli 或 Sepolia）。
//铸造 NFT 并在测试网环境中查看。

//任务步骤

//编写 NFT 合约
//使用 OpenZeppelin 的 ERC721 库编写一个 NFT 合约。
//合约应包含以下功能：
//构造函数：设置 NFT 的名称和符号。
//mintNFT 函数：允许用户铸造 NFT，并关联元数据链接（tokenURI）。
//在 Remix IDE 中编译合约。

//准备图文数据
//准备一张图片，并将其上传到 IPFS（可以使用 Pinata 或其他工具）。
//创建一个 JSON 文件，描述 NFT 的属性（如名称、描述、图片链接等）。
//将 JSON 文件上传到 IPFS，获取元数据链接。
//JSON文件参考 https://docs.opensea.io/docs/metadata-standards

//部署合约到测试网
//在 Remix IDE 中连接 MetaMask，并确保 MetaMask 连接到 Goerli 或 Sepolia 测试网。
//部署 NFT 合约到测试网，并记录合约地址。

//铸造 NFT
//使用 mintNFT 函数铸造 NFT：
//在 recipient 字段中输入你的钱包地址。
//在 tokenURI 字段中输入元数据的 IPFS 链接。
//在 MetaMask 中确认交易。

//查看 NFT
//打开 OpenSea 测试网 或 Etherscan 测试网。
//连接你的钱包，查看你铸造的 NFT。

// 作业3：编写一个讨饭合约
//任务目标
//使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。
//记录每个捐赠者的地址和捐赠金额。
//允许合约所有者提取所有捐赠的资金。

//任务步骤

//编写合约
//创建一个名为 BeggingContract 的合约。
//合约应包含以下功能：
//一个 mapping 来记录每个捐赠者的捐赠金额。
//一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
//一个 withdraw 函数，允许合约所有者提取所有资金。
//一个 getDonation 函数，允许查询某个地址的捐赠金额。
//使用 payable 修饰符和 address.transfer 实现支付和提款。

//部署合约
//在 Remix IDE 中编译合约。
//部署合约到 Goerli 或 Sepolia 测试网。

//测试合约
//使用 MetaMask 向合约发送以太币，测试 donate 功能。
//调用 withdraw 函数，测试合约所有者是否可以提取资金。
//调用 getDonation 函数，查询某个地址的捐赠金额。

//任务要求
//合约代码：
//使用 mapping 记录捐赠者的地址和金额。
//使用 payable 修饰符实现 donate 和 withdraw 函数。
//使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。

//测试网部署：
//合约必须部署到 Goerli 或 Sepolia 测试网。
//功能测试：
//确保 donate、withdraw 和 getDonation 函数正常工作。

//提交内容
//合约代码：提交 Solidity 合约文件（如 BeggingContract.sol）。
//合约地址：提交部署到测试网的合约地址。
//测试截图：提交在 Remix 或 Etherscan 上测试合约的截图。

//额外挑战（可选）
//捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
//捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
//时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。
