package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// PutStatic indexbyte is index from constanpool
type PutStatic struct {
	base.InstructionIndex16
}

func (i PutStatic) Execute(frame *runtime.Frame) {
	fieldInfo := frame.GetConstantInfo(i.Index).(*constant_pool.ConstantFieldInfo)
	utils.AssertFalse(fieldInfo == nil, "static field is null")
	_, fieldDesc := fieldInfo.NameAndDescriptor()
	// if the class is uninitiallized
	className := fieldInfo.ClassName()
	jci := klass.Perm().Get(className)
	if !jci.IsInit {
		frame.RevertPC()
		base.InitClass(jci, frame.Thread)
		return
	}
	var slots []utils.Slot
	slots = append(slots, utils.Slot{})
	if fieldDesc == "J" || fieldDesc == "D" {
		slots = append(slots, frame.PopSlot())
	}
	// new static val
	slots[0] = frame.PopSlot()
	name, _ := fieldInfo.NameAndDescriptor()

	jci.StaticFields.SetField(name, slots)

}
