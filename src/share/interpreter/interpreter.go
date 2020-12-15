package interpreter

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// code 解释器
func Interpret(method *jclass.MethodInfo) {
	var newThread = &runtime.Thread{
		PC:    1,
		Stack: runtime.NewStack(1024),
	}
	code, err := method.Attributes().AttrCode()
	if err != nil {
		return
	}
	newFrame := runtime.NewFrame(code.MaxLocals, code.MaxStack, method, newThread)
	newThread.Push(newFrame)
	loop(newThread)
}

func loop(thread *runtime.Thread) {
	reader := &base.MethodCodeReader{}
	for {
		// 因为可能在指令的操作中会对线程的栈帧进行修改，所以这个地方每次都需要进行重新赋值
		curFrame := thread.Peek()
		pc := curFrame.PC()
		thread.PC = pc
		attrCode, _ := curFrame.Method().Attributes().AttrCode()
		reader.Reset(attrCode.Code(), pc)
		opcode := reader.ReadOpenCdoe()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		curFrame.SetPC(reader.PC())
		inst.Execute(curFrame)
		if thread.IsEmtpy() {
			break
		}
	}
}
