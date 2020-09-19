package references

import (
	"../../instructions/base"
	"../../rtda"
	"../../rtda/heap"
)

/*
用于调用接口方法
*/
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

/*
invoke_interface指令需要读取32位的操作码
前2字节和其他指令一样，作为运行池索引
后两个字节中，第一个字节表示方法参数数量，第二个字节留给oracle的某些虚拟机实现使用的
*/
func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count reader.ReadUint8()
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	//fmt.Println("[gvm][invokeinterface.Execute] 执行invokeinterface命令")
	// 获取类的运行常量池
	cp := frame.Method().Class().ConstantPool()
	// 根据索引获取常量池中的接口方法符号引用
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	// 解析接口方法符号引用
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {

		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
