package native

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

func InitFloat() {
	_float(floatToRawIntBits, "floatToRawIntBits", "(F)I")
}

func _float(method Method, name, desc string) {
	Register("java/lang/Float", name, desc, method)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *runtime.Frame) {
	//value := frame.GetRef(0)
	//bits := math.Float32bits(value)

	frame.PushInt(0) // todo
}
