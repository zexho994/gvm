package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
)

type Checkcast struct {
	base.InstructionIndex16
}

func (c *Checkcast) Execute(frame *runtime.Frame) {
	ck := frame.GetConstantClassInfo(c.Index)
	k := klass.Perm.Get(ck.Name())
	ref := frame.PopRef()
	frame.PushRef(ref)

	if ref == nil || k == nil {
		return
	}

	if !ref.IsInstanceOf(k) {
		panic("checkcast error")
	}
}
