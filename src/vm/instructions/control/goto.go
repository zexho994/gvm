package control

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/rtda"
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
