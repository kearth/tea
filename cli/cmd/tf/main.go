package main

import (
	"bufio"
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"gopkg.in/yaml.v3"
)

// 嵌入的build.yaml文件
//
//go:embed build.yaml
var buildYamlFS embed.FS

// 嵌入的data.bin文件
//
//go:embed data.bin
var dataBinFS embed.FS

// BuildInfo 构建信息结构体
type BuildInfo struct {
	BuildTime       string `yaml:"build_time"`
	BuildGoVersion  string `yaml:"build_go_version"`
	BuildTeaVersion string `yaml:"build_tea_version"`
	BuildGitCommit  string `yaml:"build_git_commit"`
}

// BuildInfoData 解析yaml文件的根结构体
type BuildInfoData struct {
	BuildInfo BuildInfo `yaml:"build_info"`
}

// Constants 常量
const (
	// Unknown 未知值
	Unknown = "unknown"
	// DataBin 数据文件
	DataBin = "data.bin"
)

// Vars 全局变量
var (
	// CryptoKey 用于加密/解密data.bin文件的密钥
	CryptoKey = []byte("9f3k8m2wvq5txr7jd4hb9ep1c6n0ygsa")
	// ExePath 将在运行时设置为可执行文件的路径
	ExePath = Unknown
	// BuildGoVersion 将从build.yaml读取
	BuildGoVersion = Unknown
	// BuildTeaVersion 将从build.yaml读取，存储编译时tea库的版本
	BuildTeaVersion = Unknown
	// BuildGitCommit 将从build.yaml读取，存储构建时的git提交信息
	BuildGitCommit = Unknown
	// BuildTime 将从build.yaml读取，存储构建时间
	BuildTime = Unknown
)

// loadBuildInfo 从build.yaml文件加载构建信息
func loadBuildInfo() {

	// 从嵌入的文件系统读取build.yaml
	yamlData, err := buildYamlFS.ReadFile("build.yaml")
	if err != nil {
		fmt.Printf("警告: 无法读取嵌入的build.yaml文件: %v\n", err)
		return
	}

	// 解析yaml数据
	var buildInfo BuildInfoData
	if err := yaml.Unmarshal(yamlData, &buildInfo); err != nil {
		fmt.Printf("警告: 无法解析build.yaml文件: %v\n", err)
		return
	}

	// 更新全局变量
	BuildTime = buildInfo.BuildInfo.BuildTime
	BuildGoVersion = buildInfo.BuildInfo.BuildGoVersion
	BuildTeaVersion = buildInfo.BuildInfo.BuildTeaVersion
	BuildGitCommit = buildInfo.BuildInfo.BuildGitCommit
}

func main() {
	// 加载构建信息
	loadBuildInfo()

	// 如果没有提供参数，显示帮助信息
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	// 获取可执行文件路径
	ExePath, _ = os.Executable()
	if ExePath == "" {
		ExePath = Unknown
	}

	// 处理子命令
	cmd := os.Args[1]
	switch cmd {
	case "version":
		HandleVersionCommand()
	case "init":
		handleInitCommand()
	case "update":
		HandleUpdateCommand()
	case "run":
		HandleRunCommand()
	case "innerpack":
		PackExamples(os.Args[2])
	case "help":
		showHelp()
	default:
		fmt.Printf("未知命令: %s\n", cmd)
		showHelp()
	}
}

// ErrorPrintln 打印错误信息
func ErrorPrintln(msg string) {
	fmt.Printf("错误: %s\n", msg)
}

// HandleRunCommand 处理run命令
func HandleRunCommand() {
	// 创建子命令的flag集
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	port := runCmd.Int("port", 9106, "服务器端口（可选，默认为9106）")
	addr := runCmd.String("addr", "localhost", "服务器地址（可选，默认为'0.0.0.0'）")
	runCmd.Parse(os.Args[2:])

	// 判断当前是不是在项目根目录
	// 判断方法: 检查是否存在manifest/cmd/server/main.go文件
	runPath := filepath.Join(".", "manifest", "cmd", "server", "main.go")
	if _, err := os.Stat(runPath); os.IsNotExist(err) {
		ErrorPrintln("请在项目根目录下运行run命令!")
		os.Exit(1)
	}

	// 构建并运行服务器
	cmd := exec.Command("go", "run", runPath, fmt.Sprintf("--port=%d", *port), fmt.Sprintf("--addr=%s", *addr))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		ErrorPrintln(fmt.Sprintf("启动服务器失败 - %v", err))
		os.Exit(1)
	}

}

// 处理version命令
func HandleVersionCommand() {
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
	fmt.Printf("  Installed At: %s\n", ExePath)
	fmt.Printf("  Built Go Version: %s\n", BuildGoVersion)
	fmt.Printf("  Built Tea Version: %s\n", BuildTeaVersion)
	fmt.Printf("  Git Commit: %s\n", BuildGitCommit)
	fmt.Printf("  Built Time: %s\n", BuildTime)
}

// 打印依赖版本信息
func printDependencyVersions() string {
	var out string
	// 使用go list命令检查本地是否安装了tea模块及版本
	cmd := exec.Command("go", "list", "-m", "github.com/kearth/tea@latest")
	output, err := cmd.CombinedOutput()
	versionInfo := strings.TrimSpace(string(output))

	if err != nil || versionInfo == "" {
		versionInfo = Unknown
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

// 显示帮助信息
func showHelp() {
	fmt.Println("当前版本:", BuildTeaVersion)
	fmt.Println("")
	fmt.Println("tf工具使用说明:")
	fmt.Println("  tf version        显示工具版本号")
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
	fmt.Println("  --framework        更新tea框架")
	fmt.Println("  (默认更新tf工具)")
	fmt.Println()
	fmt.Println("run命令参数:")
	fmt.Println("  --port, -p        服务器端口（可选，默认为9106）")
	fmt.Println("  --addr, -a        服务器地址（可选，默认为'0.0.0.0'）")

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
	if destDir == "" {
		destDir = projectName
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
	if err := copyDir(destDir, modulePath); err != nil {
		// 如果复制失败，尝试清理已创建的目录
		os.RemoveAll(destDir)
		return err
	}

	return nil
}

// 复制目录内容并替换文件中的模块路径
func copyDir(dest, modulePath string) error {
	// 加载 examples 目录下的所有文件
	LoadExamples()

	files := gres.ScanDir("examples", "*", true)
	for _, file := range files {
		// 计算相对路径
		relPath, err := filepath.Rel("examples", file.Name())
		if err != nil {
			return err
		}
		// 跳过go.mod和go.sum文件
		if filepath.Base(relPath) == "go.mod" || filepath.Base(relPath) == "go.sum" {
			continue
		}

		// 跳过log目录
		if filepath.Base(relPath) == "log" || filepath.Ext(relPath) == ".log" {
			continue
		}

		// 构建目标路径
		destPath := filepath.Join(dest, relPath)

		if file.FileInfo().IsDir() {
			// 创建目标目录
			os.MkdirAll(destPath, file.FileInfo().Mode())
			continue
		}

		if filepath.Ext(destPath) == ".go" {
			reader := bufio.NewReader(bytes.NewReader(file.Content()))
			// 如果是Go文件，需要替换import语句中的"example/local"
			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()
			if err := replaceImportPaths(reader, destFile, modulePath); err != nil {
				return err
			}
		} else {
			if err = file.Export(dest, gres.ExportOption{
				RemovePrefix: "examples",
			}); err != nil {
				return fmt.Errorf("导出文件 %s 失败: %v", file.Name(), err)
			}
		}
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

// 处理update命令
func HandleUpdateCommand() {
	// 创建子命令的flag集
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	frameworkFlag := updateCmd.Bool("framework", false, "仅更新tea框架")
	// 解析update命令的参数
	updateCmd.Parse(os.Args[2:])

	// 手动检查短参数
	fmt.Println("开始更新...")

	// 更新tea框架
	if *frameworkFlag {
		fmt.Println("\n正在更新tea框架...")
		if err := updateFramework(); err != nil {
			fmt.Printf("tea框架更新失败: %v\n", err)
		} else {
			fmt.Println("tea框架更新成功!")
		}
	} else {
		// 更新tf工具
		fmt.Println("\n正在更新tf工具...")
		if err := updateTF(); err != nil {
			fmt.Printf("tf工具更新失败: %v\n", err)
		} else {
			fmt.Println("tf工具更新成功!")
		}
	}

	fmt.Println("\n更新操作完成。")
}

// 更新tf工具
func updateTF() error {
	// 验证go命令是否可用
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("无法找到go命令: %v", err)
	}

	fmt.Println("使用go install安装最新版本的tf工具...")
	// 使用go install安装最新版本
	fmt.Println("清除模块缓存...")
	cmd := exec.Command("go", "clean", "-modcache")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("清除模块缓存失败: %v\n输出: %s", err, output)
	}

	cmd = exec.Command("go", "install", "github.com/kearth/tea/cli/cmd/tf@latest")
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("安装失败: %v\n输出: %s", err, output)
	}

	return nil
}

// 更新tea框架
func updateFramework() error {
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("无法找到go命令: %v", err)
	}

	// 判断当前目录是否有 go.mod 文件
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return fmt.Errorf("当前目录没有 go.mod 文件，无法更新tea框架")
	}

	fmt.Println("清除模块缓存...")
	cmd := exec.Command("go", "clean", "-modcache")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("清除模块缓存失败: %v\n输出: %s", err, output)
	}

	// 使用go get更新tea框架到最新版本
	fmt.Println("下载最新版本的tea框架...")
	cmd = exec.Command("go", "get", "-u", "github.com/kearth/tea@latest")
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("更新失败: %v\n输出: %s", err, output)
	}

	// 执行go mod tidy确保依赖正确
	fmt.Println("更新依赖...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	output, err = tidyCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("依赖更新失败: %v\n输出: %s", err, output)
	}

	return nil
}

// PackExamples 打包 examples 目录下的所有文件
func PackExamples(path string) {
	binContent, err := gres.PackWithOption(path, gres.Option{})
	if err != nil {
		panic(err)
	}
	binContent, err = gaes.Encrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gfile.PutBytes("cli/cmd/tf/data.bin", binContent); err != nil {
		panic(err)
	}
}

// LoadExamples 加载 examples 目录下的所有文件
func LoadExamples() {
	dataBinContent, err := dataBinFS.ReadFile(DataBin)
	if err != nil {
		fmt.Printf("无法读取data.bin文件: %v\n", err)
		return
	}
	binContent, err := gaes.Decrypt(dataBinContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gres.Add(string(binContent)); err != nil {
		panic(err)
	}
}
