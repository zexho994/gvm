package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type INVOKE_VIRTUAL struct {
	base.InstructionIndex16
}

func (i *INVOKE_VIRTUAL) Execute(frame *runtime.Frame) {
	invokeClassInsance := frame.OperandStack().PopRef()
	exception.AssertTrue(invokeClassInsance != nil, "NullPointerException")
	constantMethod := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantMethod)
	name, desc := constantMethod.NameAndDescriptor()
	exception.AssertTrue(name != "<init>" && name != "<clinit>", "IncompatibleClassChangeError")
	permSpace := jclass.GetPerm().Space
	jc := permSpace[name]
	exception.AssertTrue(jc != nil, "NullPointerException")
	methodInfo, _ := jc.Methods.FindMethod(name, desc)
	if methodInfo != nil {
		exception.AssertFalse(jclass.IsStatic(methodInfo.AccessFlag()), "IncompatibleClassChangeError")

	}

}
