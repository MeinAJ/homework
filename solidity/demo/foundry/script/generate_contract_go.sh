# 创建生成脚本 generate_bindings.sh
#!/bin/bash
# 接受参数
CONTRACT=$1

out_dir="/Users/ja/project/contract/demo/foundry/out"
sol_source_dir="/Users/ja/project/contract/demo/foundry/sol_source"
go_contract_dir="/Users/ja/project/contract/demo/foundry/go_contract"

# 从编译输出中提取 ABI 和字节码
jq .abi $out_dir/"$CONTRACT".sol/"$CONTRACT".json > $sol_source_dir/"$CONTRACT".abi
jq -r .bytecode.object $out_dir/"$CONTRACT".sol/"$CONTRACT".json > $sol_source_dir/"$CONTRACT".bin

# 使用 abigen 生成 Go 代码
abigen --abi=$sol_source_dir/"$CONTRACT".abi --bin=$sol_source_dir/"$CONTRACT".bin --pkg=go_contract --type="$CONTRACT" --out=$go_contract_dir/"$CONTRACT".go