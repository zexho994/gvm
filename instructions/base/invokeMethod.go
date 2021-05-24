package base

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/native"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// InvokeMethod 执行方法调用
// 对于静态方法，方法参数就是声明的几个参数
// 对于实例方法，参数要加上编译器添加的this
// 对于本地方法，
// 对于接口方法，
func InvokeMethod(frame *runtime.Frame, method *klass.MethodInfo, isStatic bool) {
	invokerThread := frame.Thread()
	var newFrame *runtime.Frame
	var attrCode *attribute.AttrCode

	if utils.IsNative(method.AccessFlag()) {
		nativeMethod := native.FindNativeMethod(method)
		nativeMethod(frame)
		return
	}

	attrCode, _ = method.AttrCode()
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

	fmt.Printf("=== %s invoke->  %s.%s%s === \n", frame.Method().Klass().ThisClass, method.Klass().ThisClass, method.Name(), method.Descriptor())
	invokerThread.Push(newFrame)
}

// hard code to print for gvm
func gvmPrint(method *klass.MethodInfo, frame *runtime.Frame) (ok bool) {
	if method.Klass().ThisClass == "GvmOut" && method.Name() == "to" {
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
			return false
		}
	}
	return true
}
