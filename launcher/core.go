package launcher

import (
	"github.com/zouzhihao-994/gvm/classloader"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/interpreter"
	"github.com/zouzhihao-994/gvm/jclass"
)

func StartVM() {
	classloader.InitClassLoader()
	instance := jclass.ParseInstanceByClassName(config.ClassName)
	method, err := instance.FindStaticMethod("main", "([Ljava/lang/String;)V")
	if err != nil || method == nil {
		panic(err)
	}
	interpreter.Interpret(method)
}
