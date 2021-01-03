package constants

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 从运行时常量池中提取double或者long数据压入操作数栈
type LDC2_W struct {
	base.InstructionIndex16
}

func (i LDC2_W) Execute(frmae *runtime.Frame) {
	c := frmae.Method().CP().GetConstantInfo(i.Index)
	stack := frmae.OperandStack()
	switch c.(type) {
	case *constant_pool.ConstantDouble:
		double := c.(*constant_pool.ConstantDouble)
		stack.PushDouble(double.Value())
	case *constant_pool.ConstantLong:
		long := c.(*constant_pool.ConstantLong)
		stack.PushLong(long.Value())
	default:
		exception.GvmError{Msg: "ldc2_w error,unknow type "}.Throw()
	}
}
