package math

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"math"
)
import "github.com/zouzhihao-994/gvm/src/vm/rtda"

/*
REM为求余指令
*/
type DREM struct {
	base.NoOperandsInstruction
}

type FREM struct {
	base.NoOperandsInstruction
}

type IREM struct {
	base.NoOperandsInstruction
}

type LREM struct {
	base.NoOperandsInstruction
}

/*
一般取余代码为  int result = val_1 % val_2
所以对于操作数栈来说，val1 是先压入栈的，然后是val2压入栈
所以先出栈的数（val_2）在数学层面不能为零
最后结果result压入栈中
*/
func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 == 0 {
		panic("[gvm][exception] ArithmeticException : /by zero ")
	}
	result := v2 % v1
	stack.PushInt(result)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v1 == 0 {
		panic("[gvm][exception] ArithmeticExcepition : /by zero")
	}
	// go没有提供浮点数的取模运算，用Math函数代替
	result := math.Mod(v2, v1)
	stack.PushDouble(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v1 == 0 {
		panic("[gvm][execption]ArithmeticException : /by zero")
	}
	result := math.Mod(float64(v2), float64(v1))
	stack.PushFloat(float32(result))
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v1 == 0 {
		panic("[gvm][execption]ArithmeticException : /by zero")
	}
	result := v2 % v1
	stack.PushLong(result)
}
