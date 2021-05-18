package base

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// 执行方法调用
// 对于静态方法，方法参数就是声明的几个参数
// 对于实例方法，参数要加上编译器添加的this
func InvokeMethod(frame *runtime.Frame, method *klass.MethodInfo, isStatic bool) {
	invokerThread := frame.Thread()
	var newFrame *runtime.Frame
	var attrCode *attribute.Attr_Code
	if utils.IsNative(method.AccessFlag()) {
		gvmPrint(method, frame)
		return
	}
	attrCode, _ = method.Attributes().AttrCode()
	newFrame = runtime.NewFrame(attrCode.MaxLocals, attrCode.MaxStack, method, invokerThread)
	argSlotCount := int(method.ArgSlotCount())
	var n int
	if isStatic {
		if argSlotCount == 0 {
			invokerThread.Push(newFrame)
			return
		}
		n = 1
	}
	n = argSlotCount - n
	for i := n; i >= 0; i-- {
		slot := frame.OperandStack().PopSlot()
		newFrame.LocalVars().SetSlot(uint(i), slot)
	}

	invokerThread.Push(newFrame)

}

// hard code to print for gvm
func gvmPrint(method *klass.MethodInfo, frame *runtime.Frame) {
	if method.Klass().Name() == "GvmOut" && method.Name() == "to" {
		methodDesc := klass.ParseMethodDescriptor(method.Descriptor())
		switch methodDesc.Paramters()[0] {
		case "I":
			fmt.Println(frame.OperandStack().PopInt())
			break
		case "F":
			fmt.Println(frame.OperandStack().PopFloat())
			break
		case "J":
			fmt.Println(frame.OperandStack().PopLong())
			break
		case "D":
			fmt.Println(frame.OperandStack().PopDouble())
			break
		case "Z":
			fmt.Println(frame.OperandStack().PopBoolean())
			break
		case "Ljava/lang/String;":
			fmt.Println(frame.OperandStack().PopRef().JString())
		case "B":
		case "S":
			exception.GvmError{Msg: "GvmOut Error , not support byte and short types"}.Throw()
			return
		}
	}
}
