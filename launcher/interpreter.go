package launcher

import (
	"github.com/zouzhihao-994/gvm/instructions"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// Interpret code 解释器
func Interpret(method *klass.MethodInfo, t *runtime.Thread) {
	code, err := method.AttrCode()
	utils.AssertError(err, "get arrtibute code error")

	newFrame := runtime.NewFrame(code.MaxLocals, code.MaxStack, method, t)
	t.PushFrame(newFrame)

	loop(t)
}

// loop 循环执行code指令
func loop(thread *runtime.Thread) {
	methodReader := &base.MethodCodeReader{}
	for {
		curFrame := thread.PeekFrame()
		framePC := curFrame.FramePC()
		curFrame.SetThreadPC(framePC)

		methodCode, _ := curFrame.AttrCode()
		methodReader.Reset(methodCode.Code(), framePC)

		opcode := methodReader.ReadOpenCdoe()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(methodReader)
		curFrame.SetFramePC(methodReader.MethodReaderPC())

		//fmt.Printf("----%s.%s%s class exec-> %d inst----\n",
		//	curFrame.ThisClass, curFrame.MethodName(), curFrame.MethodDescriptor(), opcode)
		inst.Execute(curFrame)

		if finished(thread) {
			return
		}
	}
}

// finished 线程任务是否执行完成
func finished(thread *runtime.Thread) bool {
	if thread.IsEmtpy() {
		return true
	}
	return false
}
