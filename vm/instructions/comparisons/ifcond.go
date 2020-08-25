package comparisons

import (
	"../../instructions/base"
	"../../rtda"
)

/*
if equals
*/
type IFEQ struct{ base.BranchInstruction }

/*
if not equals
*/
type IFNE struct{ base.BranchInstruction }

/*
if （left bigger than right）
*/
type IFLT struct{ base.BranchInstruction }

/*
if(left bigger than right or equals)
*/
type IFLE struct{ base.BranchInstruction }

/*
if(left smaller than right)
*/
type IFGT struct{ base.BranchInstruction }

/*
if(left smaller than right or equals)
*/
type IFGE struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	// if val equals zero than jump
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
