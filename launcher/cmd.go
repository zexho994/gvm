package launcher

import (
	"flag"
	"fmt"
	"github.com/zouzhihao-994/gvm/classfile"
	"github.com/zouzhihao-994/gvm/interpreter"
	"github.com/zouzhihao-994/gvm/jclass"
	"os"
)

// 命令行结构体
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

// 命令行处理方法
// 对于不同的属性,设置了不同的处理方法
func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "[gvm] print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "[gvm] print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "[gvm] pring version and exit")
	flag.BoolVar(&cmd.VersionFlag, "v", false, "[gvm] pring version and exit")
	flag.StringVar(&cmd.CpOption, "classfile", "", "[gvm] classfile")
	flag.StringVar(&cmd.CpOption, "cp", "", "[gvm] class")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "[gvm] path to jre")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "[gvm] 启用详细输出")
	flag.StringVar(&cmd.Class, "class", "", "[gvm]class file name")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd
}

// 输出用法说明
func PrintUsage() {
	fmt.Printf("[gvm][usage] : %s -Xjre [jrePath] [classPath] [args...]\n", os.Args[0])
	fmt.Printf("[gvm][help] -Xjre : jrePath is the jre folder local \n" +
		"[gvm][help] -classPath : path of the class file local,is relative path based /vm\n")
}

// 通过命令行模式启动gvm
func StartGvmByCmd() {
	cmd := ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("gvm version 2.0.0")
	} else if cmd.HelpFlag {
		PrintUsage()
	}

	if cmd.XjreOption == "" {
		cmd.XjreOption = JrePath
	}

	if cmd.CpOption == "" {
		cmd.CpOption = UserClassPath
	}

	fmt.Println("start gvm -Xjre = " + cmd.XjreOption)
	fmt.Println("start gvm -cp = " + cmd.CpOption)

	startJVM(cmd.Class, cmd.XjreOption, cmd.CpOption)
}

func startJVM(className, jrePath, userClassPath string) {
	classfile.InitClassLoader(jrePath, userClassPath)
	instance := jclass.ParseInstanceByClassName(className)
	// 执行main方法
	method, err := instance.FindStaticMethod("main", "([Ljava/lang/String;)V")
	if err != nil || method == nil {
		panic(err)
	}
	interpreter.Interpret(method)
}
