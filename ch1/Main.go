package main

import "fmt"

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

func startJvm(cmd *Cmd) {
	fmt.Printf("classpath  %s \nclass : %s \nargs : %v", cmd.cpOption, cmd.class, cmd.args)
}
