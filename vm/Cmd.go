package main

import (
	"flag"
	"fmt"
	"os"
)

/*
命令行结构体
*/
type Cmd struct {
	helpFlag         bool     // 帮助命令
	versionFlag      bool     // 版本命令
	cpOption         string   // 指定路径
	class            string   // 文件名
	args             []string // 命令行的全部参数
	XjreOption       string   // 指定jre目录的位置
	verboseClassFlag bool
	verboseInstFlag  bool
}

/*
命令行处理方法
对于不同的属性,设置了不同的处理方法
*/
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "[gvm] print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "[gvm] print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "[gvm] pring version and exit")
	flag.BoolVar(&cmd.versionFlag, "v", false, "[gvm] pring version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "[gvm] classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "[gvm] class")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "[gvm]path to jre")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", true, "[gvm]启用详细输出")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

/*
输出用法说明
*/
func printUsage() {
	fmt.Printf("[gvm][usage] : %s -Xjre [jrePath] [classPath] [args...]\n", os.Args[0])
	fmt.Printf("[gvm][help] -Xjre : jrePath is the jre folder local \n" +
		"[gvm][help] -classPath : path of the class file local,is relative path based /vm\n")
}
