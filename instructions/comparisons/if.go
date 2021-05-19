package comparisons

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type IfGe struct {
	base.InstructionIndex16
}
type IfLe struct {
	base.InstructionIndex16
}
type IfEq struct {
	base.InstructionIndex16
}
type IfNe struct {
	base.InstructionIndex16
}
type IfLt struct {
	base.InstructionIndex16
}
type IfGt struct {
	base.InstructionIndex16
}
type IfNull struct {
	base.InstructionIndex16
}
type IfNonnull struct {
	base.InstructionIndex16
}

func (i IfGe) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfLe) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfEq) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfNe) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfGt) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfLt) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfNull) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopRef()
	if val == nil {
		base.Branch(frame, int(i.Index))
	}
}
func (i IfNonnull) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopRef()
	if val != nil {
		base.Branch(frame, int(i.Index))
	}
}
