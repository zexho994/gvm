package base

import (
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

func Branch(frame *runtime.Frame, offset int) {
	pc := frame.Thread().PC
	nextPC := int(pc) + offset
	frame.SetPC(uint(nextPC))
}
