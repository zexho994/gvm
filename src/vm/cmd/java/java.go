package java

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/info"
	"github.com/zouzhihao-994/gvm/src/vm/cmd"
	"github.com/zouzhihao-994/gvm/src/vm/loader"
	"github.com/zouzhihao-994/gvm/src/vm/rtda/heap"
	"strings"
)

// 通过可执行文件启动
// 使用build编译后生成可执行文件
func startByCmd() {
	// 创建一个Cmd对象赋给cmd
	c := cmd.ParseCmd()
	if c.VersionFlag {
		fmt.Printf("[gvm] gvm version is %v \n", info.GvmInfo().Version())
	} else if c.HelpFlag || c.Class == "" {
		cmd.PrintUsage()
	} else { // 启动jvm
		startJvm(c)
	}
}

/*
Jvm启动方法
*/
func startJvm(c *cmd.Cmd) {
	// 对XjreOption和cp两个字段进行解析
	// 获取classapth对象

	if c.XjreOption == "" {
		c.XjreOption = "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre"
	}

	cp := loader.Parse(c.XjreOption, c.CpOption)

	// 类加载器加载类
	// 此时cp里的3个类加载器都已经创建好了
	classLoader := heap.NewClassLoader(cp, true)

	name := "java/src/" + c.Class

	// 解析类名
	className := strings.Replace(name, ".", "/", -1)

	// 加载类,通过类的全限定名去加载类
	loadClass := classLoader.LoadClass(className)

	// 获取main方法
	mainMethod := loadClass.GetMainMethod()

	// 解释main方法
	if mainMethod != nil {
		//interpret(mainMethod, c.VerboseInstFlag, c.Args)
	} else {
		fmt.Printf("Main method not found in class %s\n", c.Class)
	}

}
