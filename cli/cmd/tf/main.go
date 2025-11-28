package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 版本号常量
const (
	Version = "v0.1.0"
)

func main() {
	// 如果没有提供参数，显示帮助信息
	if len(os.Args) == 1 {
		showHelp()
		os.Exit(1)
	}

	// 处理子命令
	cmd := os.Args[1]
	switch cmd {
	case "version":
		handleVersionCommand()
	case "init":
		handleInitCommand()
	case "help":
		showHelp()
	default:
		fmt.Printf("未知命令: %s\n", cmd)
		showHelp()
		os.Exit(1)
	}
}

// 处理version命令
func handleVersionCommand() {
	fmt.Printf("tf工具版本: %s\n", Version)
}

// 显示帮助信息
func showHelp() {
	fmt.Println("tf工具使用说明:")
	fmt.Println("  tf version       显示工具版本号")
	fmt.Println("  tf init <name>    初始化新项目")
	fmt.Println("  tf help           显示帮助信息")
	fmt.Println()
	fmt.Println("init命令参数:")
	fmt.Println("  <name>            项目名称（必需）")
	fmt.Println("  --output, -o      输出目录（可选，默认为当前目录下的项目名称）")
	fmt.Println("  --module, -m      Go模块路径（可选，默认为'example.com/' + 项目名称）")
}

// 处理init命令
func handleInitCommand() {
	// 创建子命令的flag集
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	outputFlag := initCmd.String("output", "", "输出目录")
	moduleFlag := initCmd.String("module", "", "Go模块路径")

	// 解析init命令的参数
	initCmd.Parse(os.Args[2:])

	// 检查是否提供了项目名称
	args := initCmd.Args()
	if len(args) == 0 {
		fmt.Println("错误: 缺少项目名称")
		fmt.Println("使用方式: tf init <name>")
		os.Exit(1)
	}

	projectName := args[0]

	// 处理输出目录参数
	destDir := projectName
	if *outputFlag != "" {
		destDir = *outputFlag
	}

	// 手动检查短参数 -o 和 -m
	for i, arg := range os.Args {
		if arg == "-o" && i+1 < len(os.Args) {
			destDir = os.Args[i+1]
		} else if arg == "-m" && i+1 < len(os.Args) {
			*moduleFlag = os.Args[i+1]
		}
	}

	// 处理模块路径参数
	module := "example.com/" + projectName
	if *moduleFlag != "" {
		module = *moduleFlag
	}

	// 规范化目标目录路径
	destDir, err := filepath.Abs(destDir)
	if err != nil {
		fmt.Printf("错误: 无法解析输出目录路径: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("初始化项目: %s\n", projectName)
	fmt.Printf("输出目录: %s\n", destDir)
	fmt.Printf("Go模块路径: %s\n", module)

	// 调用初始化函数（稍后实现）
	if err := initProject(projectName, destDir, module); err != nil {
		fmt.Printf("初始化项目失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("项目初始化成功!")
}

// 初始化项目
func initProject(projectName, destDir, modulePath string) error {
	// 源目录（examples目录）
	// 获取当前可执行文件的路径
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("无法获取可执行文件路径: %v", err)
	}

	// 从可执行文件路径计算到examples目录的路径
	// 可执行文件位于bin/tf，需要回到tea目录，然后进入examples
	srcDir := filepath.Join(filepath.Dir(filepath.Dir(execPath)), "examples")
	srcDir, err = filepath.Abs(srcDir)
	if err != nil {
		return fmt.Errorf("无法解析examples目录路径: %v", err)
	}

	// 检查源目录是否存在
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return fmt.Errorf("examples目录不存在: %s", srcDir)
	}

	// 检查目标目录是否已存在
	if _, err := os.Stat(destDir); err == nil {
		return fmt.Errorf("目标目录已存在: %s", destDir)
	}

	// 创建目标目录
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("创建目标目录失败: %v", err)
	}

	// 复制文件并替换模块路径
	if err := copyDir(srcDir, destDir, modulePath); err != nil {
		// 如果复制失败，尝试清理已创建的目录
		os.RemoveAll(destDir)
		return err
	}

	return nil
}

// 复制目录内容并替换文件中的模块路径
func copyDir(src, dest, modulePath string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// 跳过go.mod和go.sum文件
		if filepath.Base(path) == "go.mod" || filepath.Base(path) == "go.sum" {
			return nil
		}

		// 构建目标路径
		destPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			// 创建目标目录
			return os.MkdirAll(destPath, info.Mode())
		} else {
			// 复制文件
			if err := copyFile(path, destPath, modulePath); err != nil {
				return fmt.Errorf("复制文件 %s 失败: %v", path, err)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// 生成新的go.mod文件
	return generateGoMod(dest, modulePath)
}

// 生成新的go.mod文件并执行go mod tidy
func generateGoMod(dir, modulePath string) error {
	goModPath := filepath.Join(dir, "go.mod")
	// 移除replace指令，使用线上正式的tea包依赖
	content := fmt.Sprintf("module %s\n\ngo 1.24\n", modulePath)
	if err := os.WriteFile(goModPath, []byte(content), 0644); err != nil {
		return err
	}

	// 执行go mod tidy下载依赖
	fmt.Println("执行 go mod tidy 下载依赖...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("go mod tidy 执行失败: %v\n输出: %s\n", err, output)
		// 只打印错误，不中断流程
	}

	return nil
}

// 复制文件并在必要时替换模块路径和import语句
func copyFile(src, dest, modulePath string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 如果是go.mod文件，需要替换模块路径
	if filepath.Base(src) == "go.mod" {
		return replaceModulePath(srcFile, destFile, modulePath)
	}

	// 如果是Go文件，需要替换import语句中的"example/local"
	if filepath.Ext(src) == ".go" {
		return replaceImportPaths(srcFile, destFile, modulePath)
	}

	// 对于其他文件，直接复制内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// 复制文件权限
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	return os.Chmod(dest, srcInfo.Mode())
}

// 替换Go文件中的import路径
func replaceImportPaths(src, dest io.Reader, modulePath string) error {
	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dest.(*os.File))
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		// 替换import语句中的"example/local"为指定的modulePath
		line = strings.ReplaceAll(line, "example/local", modulePath)

		// 写入修改后的行
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// 替换go.mod文件中的模块路径
func replaceModulePath(src, dest io.Reader, modulePath string) error {
	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dest.(*os.File))
	defer writer.Flush()

	// 标记是否已经替换了模块路径
	replaced := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// 查找模块声明行并替换
		if strings.HasPrefix(trimmedLine, "module ") && !replaced {
			// 确保替换的是"example/local"或其他默认模块路径
			line = "module " + modulePath
			replaced = true
		}

		// 写入修改后的行
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
