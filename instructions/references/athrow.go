package references

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type AThrow struct {
	base.NOP
}

func (i *AThrow) Execute(frame *runtime.Frame) {
	ex := frame.PopRef()
	if ex == nil {
		exception.GvmError{Msg: exception.NullPointException}.Throw()
	}
	// todo
	frame.PopFrame()
	return
}
