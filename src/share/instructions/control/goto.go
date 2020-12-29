package control

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type GOTO struct {
	base.InstructionIndex16
}

/*
无条件的转移
*/
func (i *GOTO) Execute(frame *runtime.Frame) {
	pc := frame.Thread().PC

	nextPC := pc + uint(i.Index)

	frame.SetPC(nextPC)
}
