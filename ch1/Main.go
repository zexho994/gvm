package main

import "fmt"

/*
启动类
如果java命令不是 -version , -help,统一作为执行处理,调用startJVM()方法
*/
func main() {
	// 创建一个Cmd对象赋给cmd
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Printf(" java version is 0.0.1 ")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

/*
Jvm启动方法
*/
func startJvm(cmd *Cmd) {
	fmt.Printf("classpath  %s \nclass : %s \nargs : %v", cmd.cpOption, cmd.class, cmd.args)
}
