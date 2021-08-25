package native

import (
	"github.com/zouzhihao-994/gvm/runtime"
	"strings"
)

func isGvmNative(methodName string) bool {
	if strings.Contains(methodName, "GvmOut") {
		return true
	}

	return false
}

func InitGvmNative() {
	to(toInt, "to", "(I)V")
}

func to(method Method, name, desc string) {
	Register("GvmOut", name, desc, method)
}

func toInt(frame *runtime.Frame) {
	v := frame.PopInt()
	println(v)
}
