package main

import (
	"./classfile"
	"./instructions"
	"./instructions/base"
	"./rtda"
	"fmt"
)

/*
获取执行方法所需的局部变量表和操作数栈空间以及方法的字节码
*/
func interpret(methodInfo *classfile.MemberInfo) {
	// 获取方法属性表
	codeAttr := methodInfo.CodeAttribute()
	fmt.Printf("[gvm][interpret] main's codeAttr : %v \n", codeAttr)

	// 获取最大局部变量表
	maxLocals := codeAttr.MaxLocals()
	fmt.Printf("[gvm][interpret] main's maxLocals : %v \n", maxLocals)

	// 获取最大栈
	maxStack := codeAttr.MaxStack()
	fmt.Printf("[gvm][interpret] main's maxStack : %v \n", maxStack)

	// 获取方法表的code
	bytecode := codeAttr.Code() // 其他代码
	fmt.Printf("[gvm][interpret] main's bytecode : %v \n", bytecode)

	// 创建新的线程
	thread := rtda.NewThread()
	fmt.Println("[gvm][interpret] create new thread")

	// 初始化线程的局部变量表和最大操作数栈
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	fmt.Println("[gvm][interpret] init maxLocals and maxStack for frame")

	// 添加栈桢
	thread.PushFrame(frame)

	// 暂时没有return方法，所以用异常代替
	defer catchErr(frame)

	//
	loop(thread, bytecode)
}

/*
捕捉异常
*/
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

/*
打印局部变量表和操作数栈内容
循环执行
	---> 计算pc - 解码指令 - 执行指令
三个步骤
*/
func loop(thread *rtda.Thread, bytecode []byte) {

	// 线程弹出栈桢
	frame := thread.PopFrame()
	fmt.Println("[gvm][loop] create new thread")
	//
	reader := &base.BytecodeReader{}
	for {
		// 获取栈桢的pc指针
		pc := frame.NextPC()
		fmt.Printf("[gvm][loop] 获取pc指针%v\n", pc)
		// 设置线程的pc指针
		thread.SetPC(pc)
		// 重新开始读取指令
		reader.Reset(bytecode, pc)
		fmt.Println("[gvm][loop 准备读取指令")
		// 获取操作码
		opcode := reader.ReadUint8()
		fmt.Printf("[gvm][loop] 获取操作码 ：%v\n", opcode)
		// 解析出指令的类型
		inst := instructions.NewInstruction(opcode)
		// 执行指令的拉取操作
		inst.FetchOperands(reader)
		fmt.Println("[gvm][loop] 执行fetchOperands")
		// 设置新的PC指针
		frame.SetNextPC(reader.PC())
		fmt.Printf("[gvm][loop] frame设置新的PC指令：%v\n", reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		// 执行指令
		inst.Execute(frame)
	}
}
