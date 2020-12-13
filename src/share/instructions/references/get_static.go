package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type GET_STATIC struct {
	base.InstructionIndex16
}

func (i *GET_STATIC) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	if cp == nil {

	}

}
