package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// PutField index指向当前类的运行时常量池的索引
type PutField struct {
	base.InstructionIndex16
}

func (i *PutField) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	fieldRef := frame.Method().GetConstantInfo(i.Index).(*constant_pool.ConstantFieldInfo)
	fieldName, fieldDesc := fieldRef.NameAndDescriptor()

	var slots utils.Slots
	slots = append(slots, utils.Slot{})
	if fieldDesc == "D" || fieldDesc == "J" {
		slots = append(slots, stack.PopSlot())
	}
	slots[0] = stack.PopSlot()
	slots[0].Type = utils.TypeMapping(fieldDesc)

	objRef := stack.PopRef()
	fields := oops.FindField(fieldName, objRef.Fields(), objRef, false)
	for idx := range slots {
		fields.Slots()[idx] = slots[idx]
	}
}
