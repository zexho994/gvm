package launcher

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/loader"
	"github.com/zouzhihao-994/gvm/native"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

func StartVM() {
	GvmEnvInit()

	classFile := loader.Loading(config.ClassName)
	k := klass.ParseToKlass(&loader.ClassReader{Bytecode: classFile})
	mainMethod := mainMethod(k)
	mainThread := createMainThread()
	initClasses(mainThread)

	Interpret(mainMethod, mainThread)
}

func createMainThread() *runtime.Thread {
	mainThrad := &runtime.Thread{
		Stack: runtime.NewStack(1024),
	}
	mainThrad.SetThreadPC(0)
	return mainThrad
}

// GvmEnvInit Perform the initialization of gvm
func GvmEnvInit() {
	loader.InitClassLoader()
	klass.InitPerm()
	native.InitNativeMethod()
}

func mainMethod(k *klass.Klass) *klass.MethodInfo {
	mainMethod, err := k.FindStaticMethod("main", "([Ljava/lang/String;)V")
	utils.AssertError(err, "find main method error")
	return mainMethod
}

func initClasses(thread *runtime.Thread) {
	loadBootStrapClass()
	loadPrimitiveClasses()
	execInit(thread)
}

func loadBootStrapClass() {
	klass.ParseByClassName(config.JObjectClassName)
	klass.ParseByClassName(config.JCloneableClassName)
	klass.ParseByClassName(config.JClassClassName)
	klass.ParseByClassName(config.JStringClassName)
	klass.ParseByClassName(config.JThreadClassName)
	klass.ParseByClassName(config.JIoSerializableClassName)
}

func loadPrimitiveClasses() {
	for _, k := range klass.PrimitiveKlasses {
		klass.ParseByClassName(k.WrapperClassName)
	}
}

func execInit(thread *runtime.Thread) {
	for _, k := range klass.Perm.Space() {
		if !k.IsInit {
			base.InitClass(k, thread)
		}
	}
}
