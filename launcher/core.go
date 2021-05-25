package launcher

import (
	"github.com/zouzhihao-994/gvm/config"
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
	initBootStrapClass()
}

func mainMethod(k *klass.Klass) *klass.MethodInfo {
	mainMethod, err := k.FindStaticMethod("main", "([Ljava/lang/String;)V")
	utils.AssertError(err, "find main method error")
	return mainMethod
}

func initBootStrapClass() {
	object := klass.ParseByClassName(config.JObjectClassName)
	clone := klass.ParseByClassName(config.JCloneableClassName)
	class := klass.ParseByClassName(config.JClassClassName)
	strings := klass.ParseByClassName(config.JStringClassName)
	thread := klass.ParseByClassName(config.JThreadClassName)
	serializeble := klass.ParseByClassName(config.JIoSerializableClassName)

	klass.Perm.Save(config.JObjectClassName, object)
	klass.Perm.Save(config.JCloneableClassName, clone)
	klass.Perm.Save(config.JCloneableClassName, class)
	klass.Perm.Save(config.JCloneableClassName, strings)
	klass.Perm.Save(config.JCloneableClassName, thread)
	klass.Perm.Save(config.JCloneableClassName, serializeble)
}
