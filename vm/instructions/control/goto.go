package control

import (
	"../../instructions/base"
	"../../rtda"
)

type GOTO struct {
	base.BranchInstruction
}

/*
无条件的转移
*/
func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
