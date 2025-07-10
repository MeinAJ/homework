package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 使用全局常量定义路径（实际项目中建议使用配置变量）
const (
	artifactDir = "/Users/ja/project/contract/demo/artifacts/contracts"
	projectDir  = "/Users/ja/project/contract/demo"
	solDir      = "/Users/ja/project/contract/demo/dapp/sol_source"
	goDir       = "/Users/ja/project/contract/demo/dapp/contract"
)

// 编译所有合约（只需执行一次）
func compileContracts() error {
	if err := runScript("npx hardhat clean"); err != nil {
		return fmt.Errorf("清理失败: %w", err)
	}
	if err := runScript("npx hardhat compile"); err != nil {
		return fmt.Errorf("编译失败: %w", err)
	}
	return nil
}

// 为单个合约生成Go绑定
func generateContractGo(contractName string) error {
	// 提取并保存ABI和BIN文件
	if err := parseABI(contractName); err != nil {
		return fmt.Errorf("解析ABI失败: %w", err)
	}

	// 生成Go绑定文件
	if err := runScript(getAbiGenScript(contractName)); err != nil {
		return fmt.Errorf("生成Go代码失败: %w", err)
	}
	return nil
}

// 从编译产物中提取ABI和Bytecode
func parseABI(contractName string) error {
	jsonFile, abiFile, binFile := getConfigPath(contractName)

	jsonByte, err := os.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("读取JSON文件失败: %w", err)
	}

	var jsonData struct {
		ABI      interface{} `json:"abi"`
		Bytecode string      `json:"bytecode"`
	}
	if err := json.Unmarshal(jsonByte, &jsonData); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	// 处理ABI
	abiBytes, err := json.Marshal(jsonData.ABI)
	if err != nil {
		return fmt.Errorf("序列化ABI失败: %w", err)
	}
	if err := os.WriteFile(abiFile, abiBytes, 0644); err != nil {
		return fmt.Errorf("写入ABI文件失败: %w", err)
	}

	// 处理Bytecode（直接写入字符串内容）
	if err := os.WriteFile(binFile, []byte(jsonData.Bytecode), 0644); err != nil {
		return fmt.Errorf("写入BIN文件失败: %w", err)
	}

	return nil
}

// 执行系统命令
func runScript(script string) error {
	if err := os.Chdir(projectDir); err != nil {
		return fmt.Errorf("切换目录失败: %w", err)
	}

	args := strings.Split(script, " ")
	cmd := exec.Command(args[0], args[1:]...)

	// 捕获命令输出以便调试
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("命令执行失败: %w\n输出: %s", err, string(output))
	}

	fmt.Printf("执行成功: %s\n输出: %s\n", script, string(output))
	return nil
}

// 获取文件路径（使用filepath确保跨平台兼容性）
func getConfigPath(contractName string) (jsonPath, abiPath, binPath string) {
	jsonPath = filepath.Join(artifactDir, contractName+".sol", contractName+".json")
	abiPath = filepath.Join(solDir, contractName+".abi.json")
	binPath = filepath.Join(solDir, contractName+".bin")
	return
}

// 构造abigen命令
func getAbiGenScript(contractName string) string {
	abiFile := filepath.Join(solDir, contractName+".abi.json")
	binFile := filepath.Join(solDir, contractName+".bin")
	outFile := filepath.Join(goDir, contractName+".go")

	return fmt.Sprintf("abigen --abi=%s --bin=%s --pkg=contract --type=%s --out=%s",
		abiFile, binFile, contractName, outFile)
}

func main() {
	// 只需编译一次，然后处理所有合约
	if err := compileContracts(); err != nil {
		log.Fatalf("编译失败: %v", err)
	}

	contracts := []string{"BeggingContract", "MyToken"}
	for _, name := range contracts {
		if err := generateContractGo(name); err != nil {
			log.Printf("生成合约 %s 失败: %v", name, err)
		}
	}
}
