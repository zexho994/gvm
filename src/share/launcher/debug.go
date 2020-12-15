package launcher

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/interpreter"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
)

// 通过debug模式启动gvm
func StartGvmByDebug() {
	classfile.InitClassLoader(JrePath, UserClassPath)
	instance := jclass.ParseInstanceByClassName("classFile")
	// 执行main方法
	method, err := instance.FindStaticMethod("main")
	if err != nil || method == nil {
		panic(err)
	}
	interpreter.Interpret(method)

}
