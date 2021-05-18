package constants

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
)

// index 是当前类的运行时常量池索引，指向int，float，string或者类，方法类型，方法句柄的符号引用
type LDC struct {
	base.InstructionIndex8
}

// 如果index指向的是一个int或者float类型，那么将常量对应的数值value入栈到操作数栈中
// 如果index指向的是string，那么将字符串数值入栈
// 如果指向的是类的符号引用，解析符号引用，将Class对象的ref入栈
// 如果是方法类型或者方法句柄的符号引用，解析，然后将MethodType或者MethodHandle入栈
func (i LDC) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	c := frame.Method().CP().GetConstantInfo(uint16(i.Index))
	switch c.(type) {
	case *constant_pool.ConstantIntegerInfo:
	case *constant_pool.ConstantFloatInfo:
		float := c.(*constant_pool.ConstantFloatInfo)
		stack.PushFloat(float.Value())
	case *constant_pool.ConstantStringInfo:
		str := c.(*constant_pool.ConstantStringInfo)
		stack.PushRef(oops.NewStringOopInstance(str.String()))
	case *constant_pool.ConstantClassInfo:
	case *constant_pool.ConstantMethodInfo:
	case *constant_pool.ConstantMethodHandleInfo:
	default:
		exception.GvmError{Msg: "ldc,constant type error"}.Throw()
	}
}
