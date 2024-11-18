package bashes

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecBashFile(filePath string) {
	// 定义脚本路径
	cmd := exec.Command("sh", filePath)

	// 获取脚本输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行脚本失败: %v", err)
	}

	// 输出结果
	fmt.Printf("脚本输出:\n%s\n", output)
}