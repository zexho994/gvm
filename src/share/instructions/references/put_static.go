package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

// indexbyte is index from constanpool
type PUT_STATIC struct {
	base.InstructionIndex16
}

func (i PUT_STATIC) Execute(frame *runtime.Frame) {
	fieldInfo := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantFieldInfo)
	exception.AssertFalse(fieldInfo == nil, "static field is null")
	_, fieldDesc := fieldInfo.NameAndDescriptor()
	// if the class is uninitiallized
	className := fieldInfo.ClassName()
	jci := jclass.GetPerm().Space[className]
	if !jci.IsInit {
		frame.RevertPC()
		base.InitClass(jci, frame.Thread())
		return
	}
	var slots []utils.Slot
	slots = append(slots, utils.Slot{})
	if fieldDesc == "J" || fieldDesc == "D" {
		slots = append(slots, frame.OperandStack().PopSlot())
	}
	// new static val
	slots[0] = frame.OperandStack().PopSlot()
	name, _ := fieldInfo.NameAndDescriptor()

	jci.StaticVars.SetField(name, slots)

}
