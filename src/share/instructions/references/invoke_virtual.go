package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	jclass "github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type INVOKE_VIRTUAL struct {
	base.InstructionIndex16
}

func (i *INVOKE_VIRTUAL) Execute(frame *runtime.Frame) {
	constantMethod := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantMethod)
	methodNameStr, methodDescStr := constantMethod.NameAndDescriptor()
	exception.AssertTrue(methodNameStr != "<init>" && methodNameStr != "<clinit>", "IncompatibleClassChangeError")
	classNameStr := constantMethod.ClassName()
	permSpace := jclass.GetPerm().Space
	jc := permSpace[classNameStr]
	exception.AssertTrue(jc != nil, "NullPointerException")
	methodInfo, _ := jc.Methods.FindMethod(methodNameStr, methodDescStr)
	if methodInfo != nil {
		exception.AssertFalse(jclass.IsStatic(methodInfo.AccessFlag()), "IncompatibleClassChangeError")
		method_Descriptor := jclass.ParseMethodDescriptor(methodInfo)
		method_Descriptor.ParamteTypes()
	}

	// 参数和引用出栈
	invokeClassInsance := frame.OperandStack().PopRef()
	exception.AssertTrue(invokeClassInsance != nil, "NullPointerException")

}
