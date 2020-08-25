package comparisons

import (
	"../../instructions/base"
	"../../rtda"
)

type IF_ACMPEQ struct{ base.BranchInstruction }

type IF_ACMPNE struct{ base.BranchInstruction }

/*
弹出两个引用对象，如果相同就跳转
*/
func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}
