package main

import "./rtda/heap"
import (
	"./instructions"
	"./instructions/base"
	"./rtda"
	"fmt"
)

/*
获取执行方法所需的局部变量表和操作数栈空间以及方法的字节码
*/
func interpret(methodInfo *heap.Method, logInst bool) {

	thread := rtda.NewThread()
	fmt.Println("[gvm][interpreter.interpret] 创建新的栈桢")
	frame := thread.NewFrame(methodInfo)
	fmt.Println("[gvm][interpreter.interpret] 新栈桢 push 到 thread")
	thread.PushFrame(frame)
	fmt.Println("[gvm][interpreter.interpret] loop")
	loop(thread, logInst)
	defer catchErr(thread)
	// 获取方法属性表
	//codeAttr := methodInfo.CodeAttribute()
	//fmt.Printf("[gvm][interpret] 方法属性表 codeAttr: %v \n", codeAttr)

	// 获取最大局部变量表
	//maxLocals := codeAttr.MaxLocals()
	//fmt.Printf("[gvm][interpret] 方法局部变量表大小 maxLocals : %v \n", maxLocals)

	// 获取最大栈
	//maxStack := codeAttr.MaxStack()
	//fmt.Printf("[gvm][interpret] 方法操作数栈大小 maxStack : %v \n", maxStack)

	// 获取方法表的code
	// code的内容是指令码与指令
	//bytecode := codeAttr.Code() // 其他代码
	//fmt.Printf("[gvm][interpret] 方法Code属性 bytecode : %v \n", bytecode)

	// 创建新的线程
	//thread := rtda.NewThread()
	//fmt.Println("[gvm][interpret] 创建新线程")

	// 初始化线程的局部变量表和最大操作数栈
	//frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	//fmt.Println("[gvm][interpret] 创建栈桢，设置局部变量表和操作数栈")

	// 添加栈桢
	//thread.PushFrame(frame)
	//fmt.Println("[gvm][interpret] 方法的栈桢压入到新线程中")

	// 暂时没有return方法，所以用异常代替
	//defer catchErr(frame)

	// 执行命令
	//loop(thread, bytecode)
}

/*
捕捉异常
*/
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}

}

/*
打印局部变量表和操作数栈内容
循环执行
	---> 计算pc - 解码指令 - 执行指令
三个步骤
*/
func loop(thread *rtda.Thread, logInst bool) {
	// 字节读取器
	fmt.Printf("[gvm][interpreter.loop] ")
	reader := &base.BytecodeReader{}
	for {
		fmt.Printf("[gvm][interpreter.loop] 获取当前线程的桢桢")
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc) // decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		fmt.Printf("[gvm][interpreter.loop] 获取指令")
		inst := instructions.NewInstruction(opcode)
		fmt.Printf("[gvm][interpreter.loop] 指令fetchOperands")
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		fmt.Printf("[gvm][interpreter.loop] 指令Execute")
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
