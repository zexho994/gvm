package lang

import (
	"github.com/zouzhihao-994/gvm/src/vm/native"
	"github.com/zouzhihao-994/gvm/src/vm/oops"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *runtime.Frame) {
	this := frame.LocalVars().GetThis()
	interned := oops.InternString(this)
	frame.OperandStack().PushRef(interned)
}
