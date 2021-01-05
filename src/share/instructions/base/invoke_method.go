package base

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/attribute"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 执行方法调用
// 对于静态方法，方法参数就是声明的几个参数
// 对于实例方法，参数要加上编译器添加的this
func InvokeMethod(frame *runtime.Frame, method *jclass.MethodInfo, isStatic bool) {
	invokerThread := frame.Thread()
	var newFrame *runtime.Frame
	var attrCode *attribute.Attr_Code
	if jclass.IsNative(method.AccessFlag()) {
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
func gvmPrint(method *jclass.MethodInfo, frame *runtime.Frame) {
	if method.JClass().Name() == "GvmOut" && method.Name() == "to" {
		methodDesc := jclass.ParseMethodDescriptor(method.Descriptor())
		switch methodDesc.Paramters()[0] {
		case "I":
			fmt.Printf("gvm.out.int => %v \n", frame.OperandStack().PopInt())
			break
		case "F":
			fmt.Printf("gvm.out.float => %v \n", frame.OperandStack().PopFloat())
			break
		case "J":
			fmt.Printf("gvm.out.long => %v \n", frame.OperandStack().PopLong())
			break
		case "D":
			fmt.Printf("gvm.out.double => %v \n", frame.OperandStack().PopDouble())
			break
		case "Z":
			fmt.Printf("gvm.out.boolean => %v \n", frame.OperandStack().PopBoolean())
			break
		case "B":
		case "S":
			exception.GvmError{Msg: "GvmOut Error , not support byte and short types"}.Throw()
			return
		}
	}
}
