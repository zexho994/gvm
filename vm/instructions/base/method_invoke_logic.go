package base

import (
	"../../rtda"
	"../../rtda/heap"
)

/*
方法调用指令
对于静态方法，方法参数就是声明的几个参数
对于事例方法，处理声明的参数，还有一个编译器添加的参数this
*/
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// 获取调用栈帧的线程
	thread := invokerFrame.Thread()
	// 创建栈帧
	newFrame := thread.NewFrame(method)
	//fmt.Println("[gvm][method_invoke_logic.InvokeMethod] 新的栈桢push到线程中")
	// 栈帧推入到线程中
	thread.PushFrame(newFrame)
	//fmt.Println("[gvm][OperandStack.NewOperandStack] 获取参数数量")
	// 获取参数数量
	argSlotSlot := int(method.ArgSlotCount())

	if argSlotSlot > 0 {
		// 根据参数的数量，从操作数栈中pop对应数量的slot
		// 将pop出的slot依次放入到局部变量表中
		// 例如poo依次推出的数位 x1,x2,x3，局部变量表中存储的格式为[x3,x2,x1]
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	if method.Name() == "registerNatives" {
		thread.PopFrame()
	}

}
