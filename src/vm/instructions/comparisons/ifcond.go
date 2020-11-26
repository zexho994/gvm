package comparisons

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

/*
if equals
*/
type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	// if val equals zero than jump
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
if not equals
*/
type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
if （left bigger than right）
*/
type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
if(left bigger than right or equals)
*/
type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
if(left smaller than right)
*/
type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
if(left smaller than right or equals)
*/
type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
