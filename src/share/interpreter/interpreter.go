package interpreter

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// code 解释器
func Interpret(method *jclass.MethodInfo) {
	newThread := runtime.NewThread()
	code, err := method.Attributes().AttrCode()
	if err != nil {
		return
	}
	newFrame := runtime.NewFrame(code.MaxLocals, code.MaxStack, method)
	newThread.PushFrame(newFrame)
	loop(newThread, code.Code)
}

func loop(thread *runtime.Thread, code []byte) {
	reader := &base.MethodCodeReader{}
	for {
		curFrame := thread.Frame()
		pc := curFrame.NextPC()
		thread.SetPC(pc)
		reader.Reset(code, pc)
		opcode := reader.ReadOpenCdoe()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		curFrame.SetNextPC(reader.PC())
		inst.Execute(curFrame)
		if thread.IsStackEmpty() {
			break
		}
	}
}
