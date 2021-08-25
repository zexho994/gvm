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
	InitGvmNative()
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

	if nativeMethod, ok := registry[key]; ok {
		fmt.Printf("----find native method -> %s----\n", key)
		return nativeMethod
	}

	if method.IsRegisterNatives() || method.IsInitIDs() {
		return EmptyNative
	}
	panic("native method not found: " + key)
}
