package main

import (
	"./classfile"
	"./classpath"
	"./rtda"
	"./rtda/heap"
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
	// 获取classapth对象
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	// 类加载器加载类
	// 此时cp里的3个类加载器都已经创建好了
	classLoader := heap.NewClassLoader(cp, true)

	// 解析类名
	className := strings.Replace(cmd.class, ".", "/", -1)

	// 加载类,通过类的全限定名去加载类
	loadClass := classLoader.LoadClass(className)

	// 获取main方法
	mainMethod := loadClass.GetMainMethod()

	// 解释main方法
	if mainMethod != nil {
		interpret(mainMethod, true)
	} else {
		fmt.Printf("没有找到该类： %s \n", cmd.class)
	}

}

/*
遍历搜寻main方法
*/
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			fmt.Printf("[gvm][getMainMethod] 找到main方法 %v\n", m.Name())
			return m
		}
	}
	return nil
}

/*
测试局部局部表
*/
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

/*
测试操作数栈
*/
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

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

// 打印字节码信息
func printClassInfo(className string, cf *classfile.ClassFile) {
	fmt.Printf("========%v 字节码信息 ========\n", className)
	fmt.Printf("[gvm] JDK版本: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("[gvm] 常量池大小: %v\n", len(cf.ConstantPool()))
	fmt.Printf("[gvm] 类访问标志: 0x%x\n", cf.AccessFlags())
	fmt.Printf("[gvm] 本类名称: %v\n", cf.ClassName())
	fmt.Printf("[gvm] 父类名称: %v\n", cf.SuperClassName())
	fmt.Printf("[gvm] 接口信息: %v\n", cf.InterfaceNames())
	fmt.Printf("[gvm] 字段数量: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("[gvm] 方法或者字段名称：%s\n", f.Name())
	}
	fmt.Printf("[gvm] 类中方法数量: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("[gvm] %s\n", m.Name())
	}
	fmt.Println("========================================")

}
