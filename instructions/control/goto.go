package control

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type GOTO struct {
	base.InstructionIndex16
}

/*
无条件的转移
*/
func (i *GOTO) Execute(frame *runtime.Frame) {
	pc := frame.Thread().PC
	nextPC := uint16(pc + uint(i.Index))
	frame.SetPC(uint(nextPC))
}
