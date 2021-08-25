package base

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InvokeMethod 执行方法调用
// 对于静态方法，方法参数就是声明的几个参数
// 对于实例方法，参数要加上编译器添加的this
// 对于本地方法，
// 对于接口方法，
func InvokeMethod(frame *runtime.Frame, method *klass.MethodKlass, isStatic bool) {
	//utils.AssertTrue(method != nil, exception.NullPointException)
	if method == nil {
		return
	}

	invokerThread := frame.Thread
	var newFrame *runtime.Frame
	var attrCode *attribute.AttrCode

	attrCode, _ = method.AttrCode()
	newFrame = runtime.NewFrame(attrCode.MaxLocals, attrCode.MaxStack, method, invokerThread)
	argSlotCount := int(method.ArgSlotCount())
	var n int
	if isStatic {
		n = 1
	} else {
		n = 0
	}

	n = argSlotCount - n
	for i := n; i >= 0; i-- {
		slot := frame.PopSlot()
		newFrame.SetSlot(uint(i), slot)
	}

	if config.LogInvoke {
		fmt.Printf("---- %s invoke method ->  %s.%s%s ---- \n", frame.ThisClass, method.ThisClass, method.MethodName(), method.MethodDescriptor())
	}

	invokerThread.PushFrame(newFrame)
}
