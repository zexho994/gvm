package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
)

// 获取类的静态字段值
// index指向当前类的运行时常量池，指向对象应该是一个字段类型的符号引用
type GET_STATIC struct {
	base.InstructionIndex16
}

func (i *GET_STATIC) Execute(frame *runtime.Frame) {
	fieldRef := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantFieldInfo)

	className := fieldRef.ClassName()
	fieldName, fieldDesc := fieldRef.NameAndDescriptor()
	jci := jclass.GetPerm().Space[className]

	// 判断是否需要进行加载
	if jci == nil {
		jci = jclass.ParseInstanceByClassName(className)
		jclass.GetPerm().Space[className] = jci
		frame.RevertPC()
		base.InitClass(jci, frame.Thread())
		return
	} else if !jci.IsInit { //判断是否需要进行初始化
		frame.RevertPC()
		base.InitClass(jci, frame.Thread())
		return
	}

	field := jci.StaticVars.GetField(fieldName)
	_, _, slots := field.Fields()
	if fieldDesc == "D" || fieldDesc == "J" {
		frame.OperandStack().PushSlot(slots[0])
		slots = slots[1:]
	}
	frame.OperandStack().PushSlot(slots[0])
}
