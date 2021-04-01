package base

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

// 指令跳转
func Branch(frame *runtime.Frame, offset int) {
	pc := frame.Thread().PC
	nextPC := int(pc) + offset
	frame.SetPC(uint(nextPC))
}
