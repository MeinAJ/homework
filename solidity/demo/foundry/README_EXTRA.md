# 运行测试
```
# 运行所有测试
forge test

# 运行特定测试
forge test --match-test test_Donate

# 带详细输出
forge test -vvv

# 带Gas报告
forge test --gas-report
```

# 测试最佳实践

```
状态重置：每个测试函数都是独立运行的，setUp() 函数确保每个测试开始前都是干净状态

时间操作：使用 vm.warp() 精确控制区块时间，测试时间边界条件

权限测试：使用 vm.prank() 模拟不同地址调用合约

资金管理：使用 vm.deal() 为测试地址提供ETH

事件验证：使用 vm.expectEmit 验证事件正确触发

异常测试：使用 vm.expectRevert 验证合约在错误条件下正确回退
```
