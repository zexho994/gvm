package main

import (
	"./classfile"
	"./classpath"
	"./rtda"
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

	// 加载class 文件
	cf := loadClass(className, cp)

	// output file information
	printClassInfo(cmd.class, cf)

	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

/*
加载解析class文件
*/
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {

	// 查找className的文件
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	// output .Class info
	fmt.Printf("[gvm][loadClass] Class file data : %v", className)

	// 解析class文件
	fmt.Println("[gvm][loadClass] load class file ....")
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(className string, cf *classfile.ClassFile) {
	fmt.Printf("========%v file information========\n", className)
	fmt.Printf("[gvm] version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("[gvm] constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("[gvm] access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("[gvm] this class: %v\n", cf.ClassName())
	fmt.Printf("[gvm] super class: %v\n", cf.SuperClassName())
	fmt.Printf("[gvm] interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("[gvm] fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("[gvm] %s\n", f.Name())
	}
	fmt.Printf("[gvm] methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("[gvm] %s\n", m.Name())
	}
	fmt.Println("========================================")

}
