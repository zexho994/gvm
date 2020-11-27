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
