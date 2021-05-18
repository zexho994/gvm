package main

import (
	"flag"
	"fmt"
	"github.com/zouzhihao-994/gvm/classloader"
	"github.com/zouzhihao-994/gvm/congifuration"
	"github.com/zouzhihao-994/gvm/interpreter"
	"github.com/zouzhihao-994/gvm/jclass"
	"os"
)

func main() {
	xjre, cp, cn := GetParameters()
	startJVM(cn, xjre, cp)
}

// Cmd 命令行结构体
type Cmd struct {
	HelpFlag         bool     // 帮助命令
	VersionFlag      bool     // 版本命令
	CpOption         string   // 指定路径
	Class            string   // 文件名
	Args             []string // 命令行的全部参数
	XjreOption       string   // 指定jre目录的位置
	verboseClassFlag bool
	VerboseInstFlag  bool
}

// ParseCmd 命令行处理方法
// 对于不同的属性,设置了不同的处理方法
func ParseCmd() (cmd *Cmd) {
	cmd = &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "[gvm] print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "[gvm] print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "[gvm] pring version and exit")
	flag.BoolVar(&cmd.VersionFlag, "v", false, "[gvm] pring version and exit")
	flag.StringVar(&cmd.CpOption, "classfile", "", "[gvm] classfile")
	flag.StringVar(&cmd.CpOption, "cp", "", "[gvm] class")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "[gvm] path to jre")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "[gvm] print gvm log")
	flag.StringVar(&cmd.Class, "class", "", "[gvm] class file name")
	flag.Parse()

	return
}

// PrintUsage 输出用法说明
func PrintUsage() {
	fmt.Println("[gvm usage]:")
	fmt.Printf("\t %s -Xjre [jrePath] [classPath] [args...]\n", os.Args[0])
	fmt.Println()
	fmt.Println("[description]:")
	fmt.Printf("\t-Xjre : jrePath is the jre folder local \n" +
		"\t-classPath : path of the class file local,is relative path \n")
}

// GetParameters 通过命令行模式启动gvm
func GetParameters() (xjre, cp, cn string) {
	cmd := ParseCmd()

	// 非启动命令
	if cmd.VersionFlag {
		fmt.Println("gvm version " + congifuration.GvmVersion)
		return
	} else if cmd.HelpFlag {
		PrintUsage()
		return
	}

	// 默认值
	if cmd.XjreOption == "" {
		cmd.XjreOption = congifuration.JrePath
	}
	if cmd.CpOption == "" {
		cmd.CpOption = congifuration.UserClassPath
	}

	cn = cmd.Class
	xjre = cmd.XjreOption
	cp = cmd.CpOption

	fmt.Println("gvm -Xjre = " + xjre)
	fmt.Println("gvm -cp = " + cp)
	fmt.Println("gvm -cn = " + cn)

	return
}

// 启动
func startJVM(className, jrePath, userClassPath string) {
	classloader.InitClassLoader(jrePath, userClassPath)
	instance := jclass.ParseInstanceByClassName(className)
	method, err := instance.FindStaticMethod("main", "([Ljava/lang/String;)V")
	if err != nil || method == nil {
		panic(err)
	}
	interpreter.Interpret(method)
}
