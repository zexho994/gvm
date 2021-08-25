package native

import (
	"fmt"
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
	fmt.Printf("%v\n", val)
}

func toString(frame *runtime.Frame) {
	val := frame.LocalVars.GetRef(0)
	fmt.Printf("%v\n", val.JString())
}

func toFloat(frame *runtime.Frame) {
	val := frame.LocalVars.GetFloat(0)
	fmt.Printf("%f\n", val)
}

func toLong(frame *runtime.Frame) {
	val := frame.LocalVars.GetLong(0)
	fmt.Printf("%v\n", val)
}

func toBool(frame *runtime.Frame) {
	val := frame.LocalVars.GetBoolean(0)
	fmt.Println(val)
}

func toDouble(frame *runtime.Frame) {
	v := frame.GetDouble(0)
	fmt.Printf("%f\n", v)
}
