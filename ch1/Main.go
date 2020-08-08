package main

import (
	"../ch1/classpath"
	"fmt"
	"strings"
)

/*
启动类
如果java命令不是 -version , -help,统一作为执行处理,调用startJVM()方法
*/
func main() {
	// 创建一个Cmd对象赋给cmd
	cmd := parseCmd()
	if cmd.versionFlag { // 查询版本
		fmt.Printf("[gvm] gvm version is 0.0.1 ")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else { // 启动jvm
		startJvm(cmd)
	}
}

/*
Jvm启动方法
*/
func startJvm(cmd *Cmd) {
	fmt.Printf("[gvm][startJvm] <XjreOption> : %v , <cpOption> : %v\n", cmd.XjreOption, cmd.cpOption)
	// 对XjreOption和cp两个字段进行解析
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	// class/java/lang/Object
	fmt.Printf("[gvm][startJvm] <class> : %v\n", cmd.class)
	className := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Printf("[gvm][startJvm] <className> : %v\n", className)

	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %v\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
