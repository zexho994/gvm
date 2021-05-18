package main

import (
	"flag"
	"fmt"
	"github.com/zouzhihao-994/gvm/classloader"
	"github.com/zouzhihao-994/gvm/config"
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
	flag.StringVar(&cmd.CpOption, "classpath", "", "[gvm] classfile")
	flag.StringVar(&cmd.CpOption, "cp", "", "[gvm] class")
	flag.StringVar(&cmd.XjreOption, "xjre", "", "[gvm] path to jre")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "[gvm] print gvm log")
	flag.StringVar(&cmd.Class, "class", "", "[gvm] class file name")
	flag.Parse()

	return
}

// PrintUsage 输出用法说明
func PrintUsage() {
	fmt.Println("[gvm usage]:")
	fmt.Printf("\t%s -xjre [jrePath] -cp [classPath] -class [class name]\n", os.Args[0])
	fmt.Println()
	fmt.Println("[description]:")
	fmt.Println("\t -xjre : jrePath is the jre folder local")
	fmt.Println("\t -cp : path of the class file local,is relative path")
	fmt.Println("\t -v : print gvm version")
	fmt.Println("\t -help : print help ablout gvm")
}

// GetParameters 通过命令行模式启动gvm
func GetParameters() (xjre, cp, cn string) {
	cmd := ParseCmd()

	// 非启动命令
	if cmd.VersionFlag {
		fmt.Println("gvm version " + config.GvmVersion)
		return
	} else if cmd.HelpFlag {
		PrintUsage()
		return
	}

	// 默认值
	if cmd.XjreOption == "" {
		cmd.XjreOption = config.JrePathDefault
	}
	if cmd.CpOption == "" {
		cmd.CpOption = config.UserClassPathDefault
	}

	cn = cmd.Class
	xjre = cmd.XjreOption
	cp = cmd.CpOption

	config.JrePath = xjre
	config.ClassPath = cp
	config.ClassName = cn

	fmt.Println("gvm -Xjre = " + xjre)
	fmt.Println("gvm -cp = " + cp)
	fmt.Println("gvm -class = " + cn)

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
