package lang

import (
	"github.com/zouzhihao-994/gvm/src/vm/native"
	"github.com/zouzhihao-994/gvm/src/vm/rtda"
	"github.com/zouzhihao-994/gvm/src/vm/rtda/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
