package base

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 执行方法调用
// 对于静态方法，方法参数就是声明的几个参数
// 对于实例方法，参数要加上编译器添加的this
func InvokeMethod(invokerFrame *runtime.Frame, method *jclass.MethodInfo) {

	invokerThread := invokerFrame.Thread()
	attrCode, _ := method.Attributes().AttrCode()
	var newFrame *runtime.Frame
	if jclass.IsNatice(method.AccessFlag()) {
		method.InjectCodeAttr()
	} else {
		newFrame = runtime.NewFrame(attrCode.MaxLocals, attrCode.MaxStack, method, invokerThread)
	}
	invokerThread.Push(newFrame)
	argSlotCount := method.ArgSlotCount()
	if argSlotCount == 0 {
		return
	}
	for i := argSlotCount - 1; i >= 0; i-- {
		slot := invokerFrame.OperandStack().PopSlot()
		newFrame.LocalVars().SetSlot(i, slot)
	}
}
