package base

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InvokeMethod 执行方法调用
// 对于静态方法，方法参数就是声明的几个参数
// 对于实例方法，参数要加上编译器添加的this
// 对于本地方法
// 对于接口方法
func InvokeMethod(frame *runtime.Frame, method *klass.MethodKlass, isStatic bool) {
	if method == nil {
		return
	}

	invokerThread := frame.Thread
	attrCode, _ := method.AttrCode()
	newFrame := runtime.NewFrame(attrCode.MaxLocals, attrCode.MaxStack, method, invokerThread)

	setArguments(frame, newFrame, method, isStatic)

	if config.LogInvoke {
		fmt.Printf("---- %s invoke method ->  %s.%s%s ---- \n", frame.ThisClass, method.ThisClass, method.MethodName(), method.MethodDescriptor())
	}

	invokerThread.PushFrame(newFrame)
}

// 设置参数
func setArguments(frame *runtime.Frame, newFrame *runtime.Frame, method *klass.MethodKlass, isStatic bool) {
	if isStatic {
		setStaticArguments(frame, newFrame, method)
	} else {
		setVirtualArguments(frame, newFrame, method)
	}
}

// 设置虚方法的参数
func setVirtualArguments(frame *runtime.Frame, newFrame *runtime.Frame, method *klass.MethodKlass) {
	argSlotCount := int(method.ArgSlotCount())
	n := argSlotCount
	methodParamters := method.Descriptor().Paramters()
	for i := n; i >= 0; i-- {
		if i > 0 && methodParamters[i-1] == "D" {
			val := frame.PopDouble()
			newFrame.SetDouble(uint(i), val)
			i--
		} else if i > 0 && methodParamters[i-1] == "L" {
			val := frame.PopLong()
			newFrame.SetLong(uint(i), val)
			i--
		} else {
			slot := frame.PopSlot()
			newFrame.SetSlot(uint(i), slot)
		}
	}
}

// 设置静态方法的参数
func setStaticArguments(frame *runtime.Frame, newFrame *runtime.Frame, method *klass.MethodKlass) {
	argSlotCount := int(method.ArgSlotCount())
	n := argSlotCount - 1
	methodParamters := method.Descriptor().Paramters()
	for i := n; i >= 0; i-- {
		if methodParamters[i] == "D" {
			val := frame.PopDouble()
			newFrame.SetDouble(uint(i), val)
			i--
		} else if methodParamters[i] == "J" {
			val := frame.PopLong()
			newFrame.SetLong(uint(i), val)
			i--
		} else {
			slot := frame.PopSlot()
			newFrame.SetSlot(uint(i), slot)
		}
	}
}
