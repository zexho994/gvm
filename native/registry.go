package native

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
)

type Method func(frame *runtime.Frame)

// InitNativeMethod when invoke StartVM()
func InitNativeMethod() {
	InitVM()
	InitSystem()
	InitClassStatic()
	InitFloat()
	InitDouble()
}

var registry = map[string]Method{}

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

func FindNativeMethod(method *klass.MethodKlass) Method {
	key := method.Klass.ThisClass + "~" + method.MethodName() + "~" + method.MethodDescriptor()

	// 判断是不是gvm是实现的
	if isGvmNative(key) {

	}

	if nativeMethod, ok := registry[key]; ok {
		fmt.Printf("find native method -> %s \n", key)
		return nativeMethod
	}

	if method.IsRegisterNatives() || method.IsInitIDs() {
		return EmptyNative
	}
	panic("native method not found: " + key)
}
