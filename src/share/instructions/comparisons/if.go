package comparisons

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type IF_GE struct {
	base.InstructionIndex16
}
type IF_LE struct {
	base.InstructionIndex16
}
type IF_EQ struct {
	base.InstructionIndex16
}
type IF_NE struct {
	base.InstructionIndex16
}
type IF_LT struct {
	base.InstructionIndex16
}
type IF_GT struct {
	base.InstructionIndex16
}

func (i IF_GE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IF_LE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IF_EQ) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IF_NE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IF_GT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IF_LT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, int(i.Index))
	}
}
