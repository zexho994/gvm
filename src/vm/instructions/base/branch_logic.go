package base

import "../../rtda"

// 指令跳转
func Branch(frame *rtda.Frame, offset int) {

	pc := frame.Thread().PC()

	nextPC := pc + offset

	frame.SetNextPC(nextPC)
}
