package references

import (
	"../../rtda"
	"../../rtda/heap"
	"../base"
)

type NEW struct {
	base.Index16Instruction
}

/*
NEW指令执行
*/
func (self *NEW) Execute(frame *rtda.Frame) {
	// 获取运行时常量值
	cp := frame.Method().Class().ConstantPool()
	// 在常量池中根据index获取到类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类符号引用
	class := classRef.ResolvedClass()

	// 如果类还没有初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 如果时接口类或者抽象类，不能事例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	// 类 new一个对象引用
	ref := class.NewObject()
	// 对象引用放到操作数栈中
	frame.OperandStack().PushRef(ref)
}
