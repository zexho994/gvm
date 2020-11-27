package main

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/info"
	"github.com/zouzhihao-994/gvm/src/vm/loader"
	"github.com/zouzhihao-994/gvm/src/vm/oops"
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
	if param.jre == "" {
		param.jre = info.DefaultJrePath
	}
	if param.cn == "" {
		param.cn = "FibonacciTest"
	}

	loaders := loader.InitLoaders(info.DefaultJrePath, param.cp)
	classLoader := oops.CreateClassLoader(loaders, param.verbose == "t")

	if param.cp == "" {
		param.cp = info.DefaultCpPath + "." + param.cn
	} else {
		param.cp = param.cp + "." + param.cn
	}

	// 加载类,通过类的全限定名去加载类
	class := classLoader.LoadClass(param.cp)

	// 获取main方法
	mainMethod := class.GetMainMethod()

	if mainMethod != nil {
		// 解释main方法
		interpret(mainMethod, param.verbose == "t", param.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", param.cp)
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
	fmt.Printf("[gvm] 字段数量: %v\n", len(cf.Fields().BaseInfo()))
	for _, f := range cf.Fields().BaseInfo() {
		fmt.Printf("[gvm] 方法或者字段名称：%s\n", f.Name())
	}
	fmt.Printf("[gvm] 类中方法数量: %v\n", len(cf.Methods().BaseInfo()))
	for _, m := range cf.Methods().BaseInfo() {
		fmt.Printf("[gvm] %s\n", m.Name())
	}
	fmt.Println("========================================")

}
