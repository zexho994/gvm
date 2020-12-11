package interpreter

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	runtime "github.com/zouzhihao-994/gvm/src/share/runtime"
)

// code 解释器
func Interpret(method *jclass.MethodInfo) {
	newThread := runtime.NewThread()
	code, err := method.Attributes().AttrCode()
	if err != nil {
		return
	}
	newFrame := runtime.NewFrame(code.MaxLocals, code.MaxStack)
	newThread.PushFrame(newFrame)
	loop(newThread)
}

func loop(thread *runtime.Thread) {
	for {
		curFrame := thread.Frame()
		pc := curFrame.NextPC()
		thread.SetPC(pc)

		if thread.IsStackEmpty() {
			break
		}
	}
}
