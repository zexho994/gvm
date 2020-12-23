package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	runtime "github.com/zouzhihao-994/gvm/src/share/runtime"
)

type INVOKE_VIRTUAL struct {
	base.InstructionIndex16
}

func (i *INVOKE_VIRTUAL) Execute(frame *runtime.Frame) {
	//
	constantMethod := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantMethod)
	methodNameStr, methodDescStr := constantMethod.NameAndDescriptor()
	exception.AssertTrue(methodNameStr != "<init>" && methodNameStr != "<clinit>", "IncompatibleClassChangeError")

	classNameStr := constantMethod.ClassName()
	permSpace := jclass.GetPerm().Space
	jc := permSpace[classNameStr]
	exception.AssertTrue(jc != nil, "NullPointerException")
	methodInfo, err, _ := jc.FindMethod(methodNameStr, methodDescStr)
	exception.AssertTrue(err == nil, "no find the method of "+methodNameStr)
	exception.AssertFalse(jclass.IsStatic(methodInfo.AccessFlag()), "IncompatibleClassChangeError")

	//jclass.ParseMethodDescriptor(methodDescStr)

	//targetMethodAttrCode, _ := methodInfo.Attributes().AttrCode()
	//targetFrame := runtime.NewFrame(targetMethodAttrCode.MaxLocals, targetMethodAttrCode.MaxStack, methodInfo, frame.Thread())
	// pop params from operand_stack according the paramsType
	//frame.OperandStack().PopByParamters(paramters, targetFrame.LocalVars(), false)

	if jclass.IsProteced(methodInfo.AccessFlag()) {
		// todo if is proteced , need to judge the relation between caller and called
	}

	base.InvokeMethod(frame, methodInfo, false)
}
