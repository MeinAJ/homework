# ERC20代币

```
# 创建一个hardhat项目
npx hardhat init

# 创建.env文件
touch .env

# 配置.env文件
INFURA_PROJECT_ID=xxx
INFURA_API_KEY_SECRET=xxx
WALLET_PRIVATE_KEY=xxx

# 安装dotenv依赖
npm install dotenv --save

# 配置hardhat.config.js
require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
    solidity: "0.8.28",
    networks: {
        sepolia: {
            url: `https://sepolia.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
            accounts: [process.env.WALLET_PRIVATE_KEY] // 使用钱包私钥
        }
    }
};

# 在contracts目录下创建AJERC20.sol文件

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

    // 代币转账逻辑（仅合约内部可调用）
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

    // 授权逻辑（仅合约内部可调用）
    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "ERC20: approve from zero address");
        require(spender != address(0), "ERC20: approve to zero address");
        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount); // 触发授权事件
    }

    // 检查并扣除授权额度（仅合约内部可调用）
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

# 执行sol编译命令
npx hardhat compile

# 部署合约到sepolia网络
mkdir scripts
touch scripts/deploy.js

deploy.js内容如下：

const {ethers} = require("hardhat");

async function main() {
    // 1. 获取部署者账户
    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);

    // 2. 获取合约工厂（自动读取 artifacts 中的 ABI 和字节码）
    const ContractFactory = await ethers.getContractFactory("AJERC20");

    // 3. 部署合约（传入构造函数参数，例如 value = 100）
    const contract = await ContractFactory.deploy("AJERC20", "AJETH", 10000000);

    // 4. 等待合约部署完成
    await contract.waitForDeployment();

    // 5. 获取合约地址
    const contractAddress = await contract.getAddress();
    console.log("合约地址:", contractAddress);

    return contractAddress;
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

# 执行部署脚本
npx hardhat run scripts/deploy.js --network sepolia

# 打开metamask，连接sepolia网络，添加代币，填写代币合约地址（合约地址），即可看到代币信息。

```