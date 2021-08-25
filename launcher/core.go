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
	k := klass.ParseToKlass(loader.NewClassReader(classFile))

	mainMethod := mainMethod(k)
	mainThread := createMainThread()

	code, _ := mainMethod.AttrCode()
	newFrame := runtime.NewFrame(code.MaxLocals, code.MaxStack, mainMethod, mainThread)
	mainThread.PushFrame(newFrame)

	//initSystemProperties(mainThread)
	initClasses(mainThread)

	loop(mainThread)
}

func initSystemProperties(thread *runtime.Thread) {
	klass.ParseByClassName(config.JSystemClassName)
	sysClass := klass.Perm.Get(config.JSystemClassName)
	propsField := sysClass.GetStaticField("props", "Ljava/util/Properties;")
	_, _, _, solt := propsField.Fields()
	if len(solt) == 0 || solt[0].Ref == nil {
		thread.RevertFramePC()
		initSys, _ := sysClass.FindStaticMethod("initializeSystemClass", "()V")
		initSysFrame := runtime.NewFrame(4, 4, initSys, thread)
		thread.PushFrame(initSysFrame)
	}
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

func mainMethod(k *klass.Klass) *klass.MethodKlass {
	mainMethod, err := k.FindStaticMethod("main", "([Ljava/lang/String;)V")
	utils.AssertError(err, "find main method error")
	return mainMethod
}

func initClasses(thread *runtime.Thread) {
	loadBootStrapClass()
	loadPrimitiveClasses()

	for _, name := range BootClassNames {
		execInit(name, thread)
	}

}

func loadBootStrapClass() {
	klass.ParseByClassName(config.JObjectClassName)
	klass.ParseByClassName(config.JCloneableClassName)
	klass.ParseByClassName(config.JClassClassName)
	klass.ParseByClassName(config.JStringClassName)
	klass.ParseByClassName(config.JThreadClassName)
	klass.ParseByClassName(config.JIoSerializableClassName)
	klass.ParseByClassName(config.JThreadGroupClassName)
	klass.ParseByClassName(config.JSystemClassName)
	klass.ParseByClassName(config.JPrintStreamClassName)
}

func loadPrimitiveClasses() {
	for _, k := range klass.PrimitiveKlasses {
		klass.ParseByClassName(k.WrapperClassName)
	}
}

func execInit(name string, thread *runtime.Thread) {
	//for _, k := range klass.Perm.Space() {
	//	if !k.IsInit {
	//		base.InitClass(k, thread)
	//	}
	//}
	base.InitClass(klass.Perm.Get(name), thread)
}
