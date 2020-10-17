package references

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/rtda"
import "github.com/zouzhihao-994/gvm/src/vm/rtda/heap" // Get static field from class

type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 解析字段符号引用
	field := fieldRef.ResolvedField()
	// 获取类
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 获取描述符号
	descriptor := field.Descriptor()
	// 获取索引
	slotId := field.SlotId()
	// 获取静态变量表
	slots := class.StaticVars()
	// 获取操作数栈
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
