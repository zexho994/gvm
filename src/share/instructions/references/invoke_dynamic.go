package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用动态方法
type INVOKE_DYNAMIC struct {
	base.InstructionIndex32
}

func (i *INVOKE_DYNAMIC) Execute(frame *runtime.Frame) {
	indexByte := uint16(i.Index >> 16)
	constInvokeDynamic := frame.Method().CP().GetConstantInfo(indexByte).(*constant_pool.ConstantInvokeDynamic)
	if constInvokeDynamic == nil {

	}
}
