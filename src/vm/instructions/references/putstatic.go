package references

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
)
import "github.com/zouzhihao-994/gvm/src/vm/runtime"
import "github.com/zouzhihao-994/gvm/src/vm/oops" // Set static field in class

/**
putstatic要给静态变量赋值
*/
type PUT_STATIC struct {
	base.Index16Instruction
}

func (self PUT_STATIC) Execute(frame *runtime.Frame) {
	// 根据栈帧获取所处的方法
	currentMethod := frame.Method()
	// 根据所处方法获取当前类
	currentClass := currentMethod.Class()
	// 获取类的运行时常量池
	cp := currentClass.ConstantPool()
	// 根据索引位置获取字段的符号引用
	fieldRef := cp.GetConstant(self.Index).(*oops.FieldRef)
	// 解析字段的符号引用
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 如果不是静态变量，抛异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 如果是final修饰的，那么只能在类初始化的时候赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 获取字段的描述符
	descriptor := field.Descriptor()
	// 获取字段索引位置
	slotId := field.SlotId()
	// 获取静态变量表
	slots := class.StaticVars()
	// 获取操作数栈
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
