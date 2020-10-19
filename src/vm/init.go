package main

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/info"
	"github.com/zouzhihao-994/gvm/src/vm/classfile"
	"github.com/zouzhihao-994/gvm/src/vm/classpath"
	"github.com/zouzhihao-994/gvm/src/vm/rtda/heap"
	"strings"
)

type initParam struct {
	// jre path
	jre string
	// class path
	cp string
	// class name
	cn string
	// args
	args []string
	// 是否打印日志
	verbose string
}

// 设置参数
func receiveParam() {
	param := initParam{}

	// 为空则默认设置为/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre
	fmt.Print("jre path :")
	fmt.Scanln(&param.jre)

	// 为空不影响
	fmt.Print("class path :")
	fmt.Scanln(&param.cp)

	// 文件名
	fmt.Print("class name :")
	fmt.Scanln(&param.cn)

	// 是否打印启动日志
	fmt.Print("verbose (t/f) :")
	fmt.Scanln(&param.verbose)

	createGVM(param)
}

// 创建GVM
func createGVM(param initParam) {
	// 对XjreOption和cp两个字段进行解析
	// 获取classapth对象

	if param.jre == "" {
		param.jre = "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre"
	}

	cp := classpath.Parse(param.jre, param.cp)

	// 类加载器加载类
	// 此时cp里的3个类加载器都已经创建好了
	classLoader := heap.NewClassLoader(cp, true)

	name := ""
	if param.cp == "" {
		name = "java/src/" + param.cn
	} else {
		name = param.cp + param.cn
	}

	// 解析类名
	className := strings.Replace(name, ".", "/", -1)

	// 加载类,通过类的全限定名去加载类
	loadClass := classLoader.LoadClass(className)

	// 获取main方法
	mainMethod := loadClass.GetMainMethod()

	if mainMethod != nil {
		// 解释main方法
		interpret(mainMethod, param.verbose == "t", param.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", param.cp)
	}
}

// 通过可执行文件启动
// 使用build编译后生成可执行文件
func startByCmd() {
	// 创建一个Cmd对象赋给cmd
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("[gvm] gvm version is %v \n", info.GvmInfo().Version())
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
	// 对XjreOption和cp两个字段进行解析
	// 获取classapth对象

	if cmd.XjreOption == "" {
		cmd.XjreOption = "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre"
	}

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	// 类加载器加载类
	// 此时cp里的3个类加载器都已经创建好了
	classLoader := heap.NewClassLoader(cp, true)

	name := "java/src/" + cmd.class

	// 解析类名
	className := strings.Replace(name, ".", "/", -1)

	// 加载类,通过类的全限定名去加载类
	loadClass := classLoader.LoadClass(className)

	// 获取main方法
	mainMethod := loadClass.GetMainMethod()

	// 解释main方法

	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

/*
打印字节码信息
*/
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
