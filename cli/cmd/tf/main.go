package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// 版本号常量
const (
	Version = "v0.1.0"
)

// 构建时注入的信息
var (
	// BuildGoVersion 将在构建时通过 ldflags 注入
	BuildGoVersion = runtime.Version() // 默认使用运行时版本，构建时会被覆盖
	// BuildTeaVersion 将在构建时通过 ldflags 注入，存储编译时tea库的版本
	BuildTeaVersion = Version // 默认使用当前定义的版本，构建时会被覆盖
	// BuildGitCommit 将在构建时通过 ldflags 注入，存储构建时的git提交信息
	BuildGitCommit = "unknown" // 默认值，构建时会被覆盖
	// BuildTime 将在构建时通过 ldflags 注入，存储构建时间
	BuildTime = time.Now().Format("2006-01-02 15:04:05") // 默认使用当前时间，构建时会被覆盖
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
	case "update":
		handleUpdateCommand()
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
	// 打印欢迎信息
	fmt.Println()
	fmt.Println("Welcome to Tea Framework!")
	fmt.Println()

	// 打印环境详情
	fmt.Println("Env Detail:")
	fmt.Printf("  Go Version: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("  Tea Framework Version:%s\n", printDependencyVersions())
	fmt.Println()

	// 打印CLI详情
	fmt.Println("CLI Detail:")
	printCLIInfo()
}

// 打印依赖版本信息
func printDependencyVersions() string {
	var out string
	// 使用go list命令检查本地是否安装了tea模块及版本
	cmd := exec.Command("go", "list", "-m", "github.com/kearth/tea@latest")
	output, err := cmd.CombinedOutput()
	versionInfo := strings.TrimSpace(string(output))

	if err != nil || versionInfo == "" {
		versionInfo = "未安装"
	}

	// 解析输出，提取模块路径和版本
	parts := strings.Fields(versionInfo)
	if len(parts) >= 2 {
		out = fmt.Sprintf(" %s", parts[1])
	} else {
		out = fmt.Sprintf(" %s", versionInfo)
	}
	return out
}

// 打印CLI信息
func printCLIInfo() {
	// 获取可执行文件路径
	execPath, err := os.Executable()
	if err != nil {
		execPath = "unknown"
	}
	fmt.Printf("  Installed At: %s\n", execPath)
	fmt.Printf("  Built Go Version: %s\n", BuildGoVersion)
	fmt.Printf("  Built Tea Version: %s\n", BuildTeaVersion)
	fmt.Printf("  Git Commit: %s\n", BuildGitCommit)
	fmt.Printf("  Built Time: %s\n", BuildTime)
}

// 显示帮助信息
func showHelp() {
	fmt.Println("tf工具使用说明:")
	fmt.Println("  tf version       显示工具版本号")
	fmt.Println("  tf init <name>    初始化新项目")
	fmt.Println("  tf update         更新tf工具和框架")
	fmt.Println("  tf help           显示帮助信息")
	fmt.Println()
	fmt.Println("init命令参数:")
	fmt.Println("  <name>            项目名称（必需）")
	fmt.Println("  --output, -o      输出目录（可选，默认为当前目录下的项目名称）")
	fmt.Println("  --module, -m      Go模块路径（可选，默认为'example.com/' + 项目名称）")
	fmt.Println()
	fmt.Println("update命令参数:")
	fmt.Println("  --tf              仅更新tf工具")
	fmt.Println("  --framework       仅更新tea框架")
	fmt.Println("  (默认同时更新tf工具和tea框架)")
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

	// 构建源目录路径的多种可能位置
	// 尝试多种路径组合，确保在本地开发和go install安装时都能找到examples目录
	srcDirs := []string{
		// 开发环境路径: bin/tf -> tea/examples
		filepath.Join(filepath.Dir(filepath.Dir(execPath)), "examples"),
		// tea目录下的examples
		filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(execPath))), "tea", "examples"),
		// 相对当前工作目录的路径
		filepath.Join(".", "examples"),
		// 从用户目录开始寻找tea/examples
		filepath.Join(os.Getenv("HOME"), "zhipu", "private", "tea", "examples"),
	}

	// 查找有效的examples目录
	srcDir := ""
	for _, dir := range srcDirs {
		absDir, err := filepath.Abs(dir)
		if err == nil {
			if _, err := os.Stat(absDir); !os.IsNotExist(err) {
				srcDir = absDir
				break
			}
		}
	}

	// 如果没有找到examples目录，报错
	if srcDir == "" {
		return fmt.Errorf("无法找到examples目录，请确保安装了tea框架或者将tf工具放在正确的位置")
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

// 处理update命令
func handleUpdateCommand() {
	// 创建子命令的flag集
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	tfFlag := updateCmd.Bool("tf", false, "仅更新tf工具")
	frameworkFlag := updateCmd.Bool("framework", false, "仅更新tea框架")

	// 解析update命令的参数
	updateCmd.Parse(os.Args[2:])

	// 手动检查短参数
	tfOnly := *tfFlag
	frameworkOnly := *frameworkFlag
	for _, arg := range os.Args {
		if arg == "--tf" {
			tfOnly = true
		} else if arg == "--framework" {
			frameworkOnly = true
		}
	}

	// 如果既没有指定--tf也没有指定--framework，则同时更新两者
	if !tfOnly && !frameworkOnly {
		tfOnly = true
		frameworkOnly = true
	}

	fmt.Println("开始更新...")

	// 更新tf工具
	if tfOnly {
		fmt.Println("\n正在更新tf工具...")
		if err := updateTF(); err != nil {
			fmt.Printf("tf工具更新失败: %v\n", err)
		} else {
			fmt.Println("tf工具更新成功!")
		}
	}

	// 更新tea框架
	if frameworkOnly {
		fmt.Println("\n正在更新tea框架...")
		if err := updateFramework(); err != nil {
			fmt.Printf("tea框架更新失败: %v\n", err)
		} else {
			fmt.Println("tea框架更新成功!")
		}
	}

	fmt.Println("\n更新操作完成。")
}

// 更新tf工具
func updateTF() error {
	// 获取可执行文件路径
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("无法获取tf工具路径: %v", err)
	}

	fmt.Println("使用go install安装最新版本的tf工具...")
	// 使用go install安装最新版本
	cmd := exec.Command("go", "install", "github.com/kearth/tea/cli/cmd/tf@latest")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("安装失败: %v\n输出: %s", err, output)
	}

	// 验证go命令是否可用
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("无法找到go命令: %v", err)
	}
	
	// 获取GOPATH环境变量
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		// 如果GOPATH未设置，使用默认值
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("无法获取用户主目录: %v", err)
		}
		goPath = filepath.Join(home, "go")
	}
	
	// 构建go install安装的二进制文件路径
	installedPath := filepath.Join(goPath, "bin", "tf")
	if runtime.GOOS == "windows" {
		installedPath += ".exe"
	}

	// 检查安装的新版本文件是否存在
	if _, err := os.Stat(installedPath); os.IsNotExist(err) {
		return fmt.Errorf("安装后的tf工具文件不存在: %s", installedPath)
	}

	// 备份旧版本
	backupPath := execPath + ".bak"
	if err := os.Rename(execPath, backupPath); err != nil {
		return fmt.Errorf("无法备份当前版本: %v", err)
	}

	// 复制新版本到原位置
	newContent, err := os.ReadFile(installedPath)
	if err != nil {
		// 恢复备份
		os.Rename(backupPath, execPath)
		return fmt.Errorf("无法读取新版本: %v", err)
	}

	if err := os.WriteFile(execPath, newContent, 0755); err != nil {
		// 恢复备份
		os.Rename(backupPath, execPath)
		return fmt.Errorf("无法写入新版本: %v", err)
	}

	// 设置可执行权限
	if err := os.Chmod(execPath, 0755); err != nil {
		return fmt.Errorf("无法设置执行权限: %v", err)
	}

	// 清理备份
	os.Remove(backupPath)

	return nil
}

// 更新tea框架
func updateFramework() error {
	// 使用go get更新tea框架到最新版本
	fmt.Println("下载最新版本的tea框架...")
	cmd := exec.Command("go", "get", "-u", "github.com/kearth/tea@latest")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("更新失败: %v\n输出: %s", err, output)
	}

	// 执行go mod tidy确保依赖正确
	fmt.Println("整理依赖...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	output, err = tidyCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("依赖整理失败: %v\n输出: %s", err, output)
	}

	return nil
}
