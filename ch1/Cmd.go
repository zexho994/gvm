package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool     // 帮助命令
	versionFlag bool     // 版本命令
	cpOption    string   // 指定路径
	class       string   // 文件名
	args        []string // 命令行的全部参数
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "[zexho] print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "[zexho] print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "[zexho] pring version and exit")
	flag.BoolVar(&cmd.versionFlag, "v", false, "[zexho] pring version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "[zexho] classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "[zexho] class")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [option] class [args...]\n", os.Args[0])
}
