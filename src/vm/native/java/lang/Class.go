package lang

import "github.com/zouzhihao-994/gvm/src/vm/oops"
import (
	"github.com/zouzhihao-994/gvm/src/vm/native"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	//native.Register(jlClass, "isInterface", "()Z", isInterface)
	//native.Register(jlClass, "isPrimitive", "()Z", isPrimitive)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *runtime.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := oops.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *runtime.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*oops.Class)

	name := class.JavaName()
	nameObj := oops.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *runtime.Frame) {
	// todo
	frame.OperandStack().PushBoolean(false)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *runtime.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*oops.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *runtime.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*oops.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}
