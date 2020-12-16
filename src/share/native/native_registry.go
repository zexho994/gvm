package native

import (
	"github.com/zouzhihao-994/gvm/src/share/runtime"
	"sync"
)

type NativeMethod func(frame *runtime.Frame)

var NMM *nativeMethodsMapping
var once *sync.Once

func GetNMM() *nativeMethodsMapping {
	once.Do(func() {
		NMM = &nativeMethodsMapping{methods: make(map[string]NativeMethod, 32)}
	})
	return NMM
}

type nativeMethodsMapping struct {
	// key: class~methodName~methodDescriptor
	// val: NativeMethod
	methods map[string]NativeMethod
}

func (n nativeMethodsMapping) Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	n.methods[key] = method
}

func (n nativeMethodsMapping) FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := n.methods[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod()
	}
	return nil
}

func emptyNativeMethod() NativeMethod {
	return nil
}
