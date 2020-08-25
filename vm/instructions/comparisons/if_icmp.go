package comparisons

import (
	"../../instructions/base"
	"../../rtda"
)

type IF_ICMPEQ struct{ base.BranchInstruction }

type IF_ICMPNE struct{ base.BranchInstruction }

type IF_ICMPLT struct{ base.BranchInstruction }

type IF_ICMPLE struct{ base.BranchInstruction }

type IF_ICMPGT struct{ base.BranchInstruction }

type IF_ICMPGE struct{ base.BranchInstruction }

/*
去除两个数，如果两者不想等就跳转
*/
func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
