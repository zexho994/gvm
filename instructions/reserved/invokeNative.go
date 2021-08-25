package reserved

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/native"
	"github.com/zouzhihao-994/gvm/runtime"
)

type InvokeNative struct {
	base.NOP
}

func (*InvokeNative) Execute(frame *runtime.Frame) {
	fmt.Printf("----invoke native;%s.%s%s class exec ----\n",
		frame.ThisClass, frame.MethodName(), frame.MethodDescriptor())
	nativeMethod := native.FindNativeMethod(frame.MethodKlass)
	nativeMethod(frame)
}
