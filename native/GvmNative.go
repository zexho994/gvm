package native

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

func InitGvmNative() {
	to(toInt, "to", "(I)V")
}

func to(method Method, name, desc string) {
	Register("GvmOut", name, desc, method)
}

func toInt(frame *runtime.Frame) {
	v1 := frame.LocalVars.GetInt(0)
	println(v1)
}
