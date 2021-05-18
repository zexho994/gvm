package lang

import (
	"github.com/zouzhihao-994/gvm/native"
	"github.com/zouzhihao-994/gvm/runtime"
)

func init() {
}

func _system(method native.Method, name, desc string) {
	native.Register("java/lang/System", name, desc, method)
}

func setOut0(frame *runtime.Frame) {
}
