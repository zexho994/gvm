package reserved

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/native"
	"github.com/zouzhihao-994/gvm/runtime"
)

type InvokeNative struct {
	base.NOP
}

func (self *InvokeNative) Execute(frame *runtime.Frame) {
	nativeMethod := native.FindNativeMethod(frame.MethodInfo)
	nativeMethod(frame)
}
