package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// indexbyte is index from constanpool
type PUT_STATIC struct {
	base.InstructionIndex16
}

func (i PUT_STATIC) Execute(frame *runtime.Frame) {
	// new static val
	newVal := frame.OperandStack().PopSlot()

	fieldInfo := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantFieldRefInfo)
	exception.AssertFalse(fieldInfo == nil, "static field is null")
	name, _ := fieldInfo.NameAndDescriptor()
	// if the class is uninitiallized
	className := fieldInfo.ClassName()
	jci := jclass.GetPerm().Space[className]
	if !jci.IsInit {
		base.InitClass(jci, frame.Thread())
	}
	jci.StaticVars.SetField(name, &newVal)

}
