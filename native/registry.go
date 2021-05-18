package native

import (
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/runtime"
)

type Method func(frame *runtime.Frame)

var registry = map[string]Method{}

func emptyNativeMethod(frame *runtime.Frame) {
	// do nothing
}
func Register(className, methodName, methodDescriptor string, method Method) {
	key := className + "~" + methodName + "~" + methodDescriptor
	if _, ok := registry[key]; !ok {
		registry[key] = method
	} else {
		panic("native method: " + key + "has been registerd!")
	}
}

func FindNativeMethod(method *jclass.MethodInfo) Method {
	key := method.JClass().Name() + "~" + method.Name() + "~" + method.Descriptor()
	if nativeMethod, ok := registry[key]; ok {
		return nativeMethod
	}

	if method.IsRegisterNatives() || method.IsInitIDs() {
		return emptyNativeMethod
	}
	panic("native method not found: " + key)
}
