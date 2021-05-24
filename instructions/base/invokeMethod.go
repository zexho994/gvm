package base

import (
	"fmt"
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
	invokerThread := frame.Thread
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
		slot := frame.PopSlot()
		newFrame.SetSlot(uint(i), slot)
	}

	fmt.Printf("=== %s invoke->  %s.%s%s === \n", frame.Method().ThisClass, method.ThisClass, method.Name(), method.Descriptor())
	invokerThread.Push(newFrame)
}
