package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

// 获取对象的字段值
type GET_FIELD struct {
	// 指向运行时常量池中字段的符号引用
	base.InstructionIndex16
}

func (i *GET_FIELD) Execute(frame *runtime.Frame) {
	objRef := frame.OperandStack().PopRef()
	exception.AssertFalse(objRef == nil, exception.NULL_POINT_EXCEPTION)

	constFieldRef := objRef.Jclass().ConstantPool.GetConstantInfo(i.Index).(*constant_pool.ConstantFieldRefInfo)
	fieldName, _ := constFieldRef.NameAndDescriptor()
	field, r := objRef.FindField(fieldName)
	exception.AssertTrue(r, exception.FIELDS_NOT_FOUND_ERROR)
	exception.AssertFalse(jclass.IsStatic(field.AccessFlag()), exception.INCOMPATIBLE_CLASS_CHANGE_ERROR)

	fieldsSlot := field.Slots()[0]
	if fieldsSlot.Type == utils.SlotLong {
		v1, v2 := field.Slots().GetVal64()
		v := int64(v1)<<32 + int64(v2)
		frame.OperandStack().PushLong(v)
		return
	}
}
