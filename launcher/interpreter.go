package launcher

import (
	"fmt"
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
	t.Push(newFrame)

	loop(t)
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
		fmt.Printf("----%s.%s%s class exec-> %d inst----\n",
			curFrame.Method().Klass().ThisClass, curFrame.Method().Name(), curFrame.Method().Descriptor(), opcode)
		inst.Execute(curFrame)
		if thread.IsEmtpy() {
			break
		}
	}
}
