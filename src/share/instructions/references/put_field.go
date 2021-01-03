package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/oops"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// index指向当前类的运行时常量池的索引
type PUT_FIELD struct {
	base.InstructionIndex16
}

func (i *PUT_FIELD) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	stack := frame.OperandStack()
	fieldRef := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantFieldRefInfo)
	if stack == nil || fieldRef == nil {

	}
	name, desc := fieldRef.NameAndDescriptor()
	switch desc {
	case "D":
		d := stack.PopDouble()
		objRef := stack.PopRef()
		fields := objRef.Fields()
		slots := oops.FindField(name, fields, objRef, false)
		slots.Slots().SetVal64(int32(d), int32(int64(d)>>32))
	case "J":
		l := stack.PopLong()
		objRef := stack.PopRef()
		fields := objRef.Fields()
		slots := oops.FindField(name, fields, objRef, false)
		slots.Slots().SetVal64(int32(l), int32(l>>32))
	default:
		slot := stack.PopSlot()
		objRef := stack.PopRef()
		slots := oops.FindField(name, objRef.Fields(), objRef, false)
		slots.Slots()[0] = slot
	}
}
