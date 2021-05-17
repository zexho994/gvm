package launcher

import (
	"github.com/zouzhihao-994/gvm/classfile"
	"github.com/zouzhihao-994/gvm/interpreter"
	"github.com/zouzhihao-994/gvm/jclass"
)

// 通过debug模式启动gvm
func StartGvmByDebug(className, jrePath, userClassPath string) {
	classfile.InitClassLoader(jrePath, userClassPath)
	instance := jclass.ParseInstanceByClassName(className)
	// 执行main方法
	method, err := instance.FindStaticMethod("main", "([Ljava/lang/String;)V")
	if err != nil || method == nil {
		panic(err)
	}
	interpreter.Interpret(method)
}