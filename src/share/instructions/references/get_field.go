package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
	"github.com/zouzhihao-994/gvm/src/share/utils"
	"math"
)

// 获取对象的字段值
type GET_FIELD struct {
	// 指向运行时常量池中字段的符号引用
	base.InstructionIndex16
}

func (i *GET_FIELD) Execute(frame *runtime.Frame) {
	objRef := frame.OperandStack().PopRef()
	exception.AssertFalse(objRef == nil, exception.NullPointException)

	constFieldRef := objRef.Jclass().ConstantPool.GetConstantInfo(i.Index).(*constant_pool.ConstantFieldRefInfo)
	fieldName, _ := constFieldRef.NameAndDescriptor()
	field, r := objRef.FindField(fieldName)
	exception.AssertTrue(r, exception.FieldsNotFoundError)
	exception.AssertFalse(jclass.IsStatic(field.AccessFlag()), exception.IncompatibleClassChangeError)

	fieldsSlot := field.Slots()[0]
	if fieldsSlot.Type == utils.SlotLong {
		v1, v2 := field.Slots().GetVal64()
		v := int64(v2)<<32 + int64(v1)
		frame.OperandStack().PushLong(v)
		return
	}
	if fieldsSlot.Type == utils.SlotDouble {
		v1, v2 := field.Slots().GetVal64()
		v := uint64(v2)<<32 + uint64(v1)
		frame.OperandStack().PushDouble(math.Float64frombits(v))
		return
	}

	frame.OperandStack().PushSlot(field.Slots()[0])
}
