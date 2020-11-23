package main

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/vm/instructions"
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/rtda"
	"github.com/zouzhihao-994/gvm/src/vm/rtda/heap"
)

/*
获取执行方法所需的局部变量表和操作数栈空间以及方法的字节码
*/
func interpret(methodInfo *heap.Method, logInst bool, args []string) {
	// 创建一个新的线程
	thread := rtda.NewThread()

	// 创建一个栈桢
	frame := thread.NewFrame(methodInfo)

	// 栈桢压入虚拟机栈
	thread.PushFrame(frame)

	jArgs := createArgsArray(methodInfo.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)

	defer catchErr(thread)

	// 执行code命令
	loop(thread, logInst)

}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
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
执行Code命令
循环执行：计算pc -> 解码指令 -> 执行指令 三个步骤
为什么说线程栈是线程私有的？
*/
func loop(thread *rtda.Thread, logInst bool) {
	// 字节读取器
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()

		// 获取栈桢的pc指针
		pc := frame.NextPC()

		// 线程设置pc指针
		thread.SetPC(pc)

		// 设置code和pc
		reader.Reset(frame.Method().Code(), pc)

		// 获取操作码，同时pc++
		opcode := reader.ReadUint8()

		//  根据操作码获取对应的命令
		//fmt.Printf("[gvm][interpreter.loop] 获取指令 \n")
		inst := instructions.NewInstruction(opcode)

		// 拉取操作数
		//fmt.Println("[gvm][interpreter.loop] 指令fetchOperands")
		inst.FetchOperands(reader)

		// 更新栈桢
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// 执行指针
		inst.Execute(frame)

		// 如果线程栈空了，就推出
		if thread.IsStackEmpty() {
			break
		}
	}
}

/*
打印指令信息
*/
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("[gvm][logInstruction] %v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
