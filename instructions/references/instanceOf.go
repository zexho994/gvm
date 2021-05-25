package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
)

type InstanceOf struct {
	base.InstructionIndex16
}

func (i *InstanceOf) Execute(frame *runtime.Frame) {
	ref := frame.PopRef()
	constClass := frame.GetConstantClassInfo(i.Index)
	t := klass.Perm.Get(constClass.Name())
	if t == nil {
		t = klass.ParseByClassName(constClass.Name())
	}

	if ref == nil {
		frame.PushInt(0)
	} else if ref.IsInstanceOf(t) {
		frame.PushInt(1)
	} else {
		frame.PushInt(0)
	}
}
