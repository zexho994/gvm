package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type INVOKE_VIRTUAL struct {
	base.InstructionIndex16
}

func (i *INVOKE_VIRTUAL) Execute(frame *runtime.Frame) {

}
