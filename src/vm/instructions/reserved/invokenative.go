package reserved

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/rtda"
import "github.com/zouzhihao-994/gvm/src/vm/native"
import _ "github.com/zouzhihao-994/gvm/src/vm/native/java/lang"
import _ "github.com/zouzhihao-994/gvm/src/vm/native/sun/misc"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
