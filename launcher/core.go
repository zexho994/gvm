package launcher

import (
	"github.com/zouzhihao-994/gvm/classloader"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

func StartVM() {
	classloader.InitClassLoader()
	instance := jclass.ParseInstanceByClassName(config.ClassName)

	mainMethod, err := instance.FindStaticMethod("main", "([Ljava/lang/String;)V")
	utils.AssertError(err, "start vm error")
	utils.AssertTrue(mainMethod != nil, "mainMethod() missing")

	mainThread := createMainThread()
	Interpret(mainMethod, mainThread)
}

func createMainThread() *runtime.Thread {
	return &runtime.Thread{
		PC:    0,
		Stack: runtime.NewStack(1024),
	}
}
