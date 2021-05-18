package launcher

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/loader"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

func StartVM() {
	loader.InitClassLoader()
	classFile := loader.Loading(config.ClassName)
	reader := &loader.ClassReader{Bytecode: classFile}
	k := klass.ParseToKlass(reader)

	mainMethod, err := k.FindStaticMethod("main", "([Ljava/lang/String;)V")
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
