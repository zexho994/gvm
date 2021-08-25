package native

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

func InitGvmNative() {
	to(toInt, "to", "(I)V")
	to(toLong, "to", "(J)V")
	to(toFloat, "to", "(F)V")
	to(toDouble, "to", "(D)V")
	to(toBool, "to", "(Z)V")
	to(toString, "to", "(Ljava/lang/String;)V")
}

func to(method Method, name, desc string) {
	Register("GvmOut", name, desc, method)
}

func toInt(frame *runtime.Frame) {
	val := frame.LocalVars.GetInt(0)
	println(val)
}

func toString(frame *runtime.Frame) {
	val := frame.LocalVars.GetRef(0)
	println(val.JString())
}

func toFloat(frame *runtime.Frame) {
	val := frame.LocalVars.GetFloat(0)
	println(val)
}

func toLong(frame *runtime.Frame) {
	val := frame.LocalVars.GetLong(0)
	println(val)
}

func toBool(frame *runtime.Frame) {
	val := frame.LocalVars.GetBoolean(0)
	println(val)
}

func toDouble(frame *runtime.Frame) {
	v := frame.GetDouble(0)

	//v := (val1 >> 16) + val2
	println(v)
}
