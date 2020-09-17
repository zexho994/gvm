package references

import (
	"../../instructions/base"
	"../../rtda"
	"../../rtda/heap"
	"fmt"
)

/*
调用静态方法
*/
type INVOKE_STATIC struct{ base.Index16Instruction }

/*

 */
func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	fmt.Println("[gvm][invokestatic.Execute] 执行invokestatic命令")
	// 运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 获取方法符号引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	// 解析符号引用 -> 获取方法结构体
	resolvedMethod := methodRef.ResolvedMethod()
	// 不是静态方法则发生异常
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 调用方法
	base.InvokeMethod(frame, resolvedMethod)
}
