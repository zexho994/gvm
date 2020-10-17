package base

import "github.com/zouzhihao-994/gvm/src/vm/rtda"

// 指令跳转
func Branch(frame *rtda.Frame, offset int) {

	pc := frame.Thread().PC()

	nextPC := pc + offset

	frame.SetNextPC(nextPC)
}