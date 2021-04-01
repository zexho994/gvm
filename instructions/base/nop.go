package base

import "github.com/zouzhihao-994/gvm/runtime"

type NOP struct {
	InstructionIndex0
}

func (NOP) Execute(*runtime.Frame) {
	// nothing
}
