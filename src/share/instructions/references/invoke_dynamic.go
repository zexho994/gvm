package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用动态方法
type INVOKE_DYNAMIC struct {
	base.InstructionIndex16
}

func (i INVOKE_DYNAMIC) Execute(frame *runtime.Frame) {
	constInvokeDynamic := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantInvokeDynamic)
	if constInvokeDynamic == nil {

	}
}
