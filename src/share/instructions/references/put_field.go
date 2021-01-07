package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/oops"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

// index指向当前类的运行时常量池的索引
type PUT_FIELD struct {
	base.InstructionIndex16
}

func (i *PUT_FIELD) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	stack := frame.OperandStack()
	fieldRef := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantFieldRefInfo)
	fieldName, fieldDesc := fieldRef.NameAndDescriptor()

	var slots utils.Slots
	slots = append(slots, utils.Slot{})
	if fieldDesc == "D" || fieldDesc == "J" {
		slots = append(slots, stack.PopSlot())
	}
	slots[0] = stack.PopSlot()
	objRef := stack.PopRef()
	fields := oops.FindField(fieldName, objRef.Fields(), objRef, false)
	for idx := range slots {
		fields.Slots()[idx] = slots[idx]
	}
}
