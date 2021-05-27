package native

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

func InitDouble() {
	_double(doubleToRawLongBits, "doubleToRawLongBits", "(D)J")
	_double(longBitsToDouble, "longBitsToDouble", "(J)D")
}

func _double(method Method, name, desc string) {
	Register("java/lang/Double", name, desc, method)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *runtime.Frame) {
	// todo
	frame.PushLong(int64(frame.GetDouble(0)))
}

func longBitsToDouble(frame *runtime.Frame) {
	frame.PushDouble(float64(frame.GetLong(0)))
}
