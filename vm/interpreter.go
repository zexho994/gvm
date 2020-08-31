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
	// 获取最大局部变量表
	maxLocals := codeAttr.MaxLocals()
	// 获取最大栈
	maxStack := codeAttr.MaxStack()
	// 获取方法表的code
	bytecode := codeAttr.Code() // 其他代码
	// 创建新的线程
	thread := rtda.NewThread()
	// 初始化线程的局部变量表和最大操作数栈
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	// 添加栈桢
	thread.PushFrame(frame)
	// 暂时没有return方法，所以用异常代替
	defer catchErr(frame)
	//
	loop(thread, bytecode)
}

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
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
	/*
		线程弹出栈桢

	*/
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc) // decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
