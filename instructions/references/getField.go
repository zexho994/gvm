package references

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// GetField 获取对象的字段值
type GetField struct {
	// 指向运行时常量池中字段的符号引用
	base.InstructionIndex16
}

func (i *GetField) Execute(frame *runtime.Frame) {
	fieldRef := frame.GetConstantFieldsInfo(i.Index)
	fieldName, fieldDesc := fieldRef.NameAndDescriptor()
	k := klass.Perm.Get(frame.ThisClass)

	objRef := frame.PopRef()
	utils.AssertFalse(objRef == nil, exception.NullPointException)
	fields := oops.FindField(fieldName, objRef.OopFields, k)

	frame.PushSlot(fields.Slots()[0])
	if fieldDesc == "D" || fieldDesc == "J" {
		frame.PushSlot(fields.Slots()[1])
	}
}
