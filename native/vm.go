package native

import "github.com/zouzhihao-994/gvm/runtime"

func InitVM() {
	_vm(initialize, "initialize", "()V")
}

func _vm(method Method, name, desc string) {
	Register("sun/misc/VM", name, desc, method)
}

func initialize(frame *runtime.Frame) {
	// todo
}
