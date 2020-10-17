package conversions

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/rtda"

/*
double to float
*/
type I2F struct{ base.NoOperandsInstruction }

/*
double to int
*/
type I2D struct{ base.NoOperandsInstruction }

/*
double to long
*/
type I2L struct{ base.NoOperandsInstruction }

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	i := float64(d)
	stack.PushDouble(i)
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	i := float32(d)
	stack.PushFloat(i)
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := int64(d)
	stack.PushLong(i)
}
