package base

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

// Branch 指令跳转
func Branch(frame *runtime.Frame, offset int) {
	pc := frame.ThreadPC()
	nextPC := int(pc) + offset
	frame.SetPC(uint(nextPC))
}
