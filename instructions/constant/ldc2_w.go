package constants

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
)

// Ldc2W 从运行时常量池中提取double或者long数据压入操作数栈
type Ldc2W struct {
	base.InstructionIndex16
}

func (i Ldc2W) Execute(frmae *runtime.Frame) {
	c := frmae.GetConstantInfo(i.Index)
	switch c.(type) {
	case *constant_pool.ConstantDoubleInfo:
		double := c.(*constant_pool.ConstantDoubleInfo)
		frmae.PushDouble(double.Value())
	case *constant_pool.ConstantLongInfo:
		long := c.(*constant_pool.ConstantLongInfo)
		frmae.PushLong(long.Value())
	default:
		exception.GvmError{Msg: "ldc2_w error,unknow type "}.Throw()
	}
}
