package native

import "github.com/zouzhihao-994/gvm/src/vm/runtime"

/*
本地方法
*/
type NativeMethod func(frame *runtime.Frame)

var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *runtime.Frame) {
	// do nothing
}

/*
方法注册
方法的key为"类名 + 方法名 + 方法描述"
*/
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

/*
查找方法

*/
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}
