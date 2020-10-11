package conversions

import "../base"
import "../../rtda"

/*
double to float
*/
type F2D struct{ base.NoOperandsInstruction }

/*
double to int
*/
type F2I struct{ base.NoOperandsInstruction }

/*
double to long
*/
type F2L struct{ base.NoOperandsInstruction }

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int32(d)
	stack.PushInt(i)
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int64(d)
	stack.PushLong(i)
}

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := float32(d)
	stack.PushFloat(i)
}
