package native

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

type Method func(frame *runtime.Frame)

var registry = map[string]Method{}

func initializeSystemClass(frame *runtime.Frame) {
	sys := klass.Perm().Space["java/lang/System"]
	if sys == nil {
		return
	}
	initSysClass, err := sys.FindStaticMethod("initializeSystemClass", "()V")
	utils.AssertError(err, "")
	newFrame := runtime.NewFrame(4, 3, initSysClass, frame.Thread())
	frame.Thread().Push(newFrame)
}

func EmptyNative(frame *runtime.Frame) {
	//
}

func Register(className, methodName, methodDescriptor string, method Method) {
	key := className + "~" + methodName + "~" + methodDescriptor
	if _, ok := registry[key]; !ok {
		registry[key] = method
	} else {
		panic("native method: " + key + "has been registerd!")
	}
}

func FindNativeMethod(method *klass.MethodInfo) Method {
	key := method.Klass().ThisClass + "~" + method.Name() + "~" + method.Descriptor()
	if nativeMethod, ok := registry[key]; ok {
		return nativeMethod
	}

	if method.IsRegisterNatives() {
		return initializeSystemClass
	}
	if method.IsInitIDs() {
		return EmptyNative
	}
	panic("native method not found: " + key)
}
