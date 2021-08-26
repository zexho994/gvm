package launcher

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/instructions"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

// loop 循环执行code指令
func loop(thread *runtime.Thread) {
	methodReader := &base.MethodCodeReader{}
	for {
		updatePC(thread)
		updateCodeArrt(thread, methodReader)
		execInst(thread, methodReader)
		if finished(thread) {
			break
		}
	}
}

// 获取thread栈顶的frame的pc -> thread的pc更新为frame的pc ->
func updatePC(thread *runtime.Thread) {
	curFrame := thread.PeekFrame()
	framePC := curFrame.FramePC()
	curFrame.SetThreadPC(framePC)
}

//更新属性表code
func updateCodeArrt(thread *runtime.Thread, reader *base.MethodCodeReader) {
	curFrame := thread.PeekFrame()
	methodCode, _ := curFrame.AttrCode()
	reader.Reset(methodCode.Code(), curFrame.FramePC())
}

//执行指令
func execInst(thread *runtime.Thread, reader *base.MethodCodeReader) {
	curFrame := thread.PeekFrame()
	//获取操作码
	opcode := reader.ReadOpenCdoe()
	//创建指令
	inst := instructions.NewInstruction(opcode)
	//获取操作数
	inst.FetchOperands(reader)
	curFrame.SetFramePC(reader.MethodReaderPC())
	if config.LogInterpreter {
		fmt.Printf("----[interpreter] %s.%s%s class exec -> %d inst----\n", curFrame.ThisClass, curFrame.MethodName(), curFrame.MethodDescriptor(), opcode)
	}
	//执行指令
	inst.Execute(curFrame)
}

// finished 线程任务是否执行完成
func finished(thread *runtime.Thread) bool {
	if thread.IsEmtpy() {
		return true
	}
	return false
}
