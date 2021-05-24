package main

import (
	"flag"
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/launcher"
	"os"
)

func main() {
	initParamConfig()
	launcher.StartVM()
}

// Cmd 命令行结构体
type Cmd struct {
	HelpFlag    bool     // 帮助命令
	VersionFlag bool     // 版本命令
	CpOption    string   // 指定路径
	Class       string   // 文件名
	Args        []string // 命令行的全部参数
	XjreOption  string   // 指定jre目录的位置
}

// ParseCmd 命令行处理方法
// 对于不同的属性,设置了不同的处理方法
func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "[gvm] print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "[gvm] print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "[gvm] pring version and exit")
	flag.BoolVar(&cmd.VersionFlag, "v", false, "[gvm] pring version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "[gvm] classfile")
	flag.StringVar(&cmd.CpOption, "cp", "", "[gvm] class")
	flag.StringVar(&cmd.XjreOption, "xjre", "", "[gvm] path to jre")
	flag.StringVar(&cmd.Class, "class", "", "[gvm] class file name")
	flag.Parse()

	return cmd
}

// PrintUsage 输出用法说明
func PrintUsage() {
	fmt.Println("[gvm usage]:")
	fmt.Printf("\t%s -xjre [jrePath] -cp [classPath] -class [class name]\n\n", os.Args[0])
	fmt.Println("[description]:")
	fmt.Println("\t -xjre : jrePath is the jre folder local")
	fmt.Println("\t -cp : path of the class file local,is relative path")
	fmt.Println("\t -v : print gvm version")
	fmt.Println("\t -help : print help ablout gvm")
}

// initParamConfig 通过命令行模式启动gvm
func initParamConfig() {
	cmd := ParseCmd()

	// 非启动命令
	if cmd.isHelpOrVersion() {
		return
	}
	cmd.checkDefault()
	cmd.updateConfig()
	cmd.printArguments()
}

func (cmd *Cmd) isHelpOrVersion() bool {
	if cmd.VersionFlag {
		fmt.Println("gvm version " + config.GvmVersion)
		return true
	} else if cmd.HelpFlag {
		PrintUsage()
		return true
	}
	return false
}

func (cmd *Cmd) checkDefault() {
	if cmd.XjreOption == "" {
		cmd.XjreOption = config.JrePathDefault
	}
	if cmd.CpOption == "" {
		cmd.CpOption = config.UserClassPathDefault
	}
}

func (cmd *Cmd) printArguments() {
	fmt.Println("gvm -Xjre = " + config.JrePath)
	fmt.Println("gvm -cp = " + config.ClassPath)
	fmt.Println("gvm -class = " + config.ClassName)
}

func (cmd *Cmd) updateConfig() {
	config.JrePath = cmd.XjreOption
	config.ClassPath = cmd.CpOption
	config.ClassName = cmd.Class
}
