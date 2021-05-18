package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

type INVOKE_VIRTUAL struct {
	base.InstructionIndex16
}

func (i *INVOKE_VIRTUAL) Execute(frame *runtime.Frame) {
	constantMethod := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantMethodInfo)
	methodNameStr, methodDescStr := constantMethod.NameAndDescriptor()
	utils.AssertTrue(methodNameStr != "<init>" && methodNameStr != "<clinit>", "IncompatibleClassChangeError")

	classNameStr := constantMethod.ClassName()
	permSpace := jclass.Perm().Space
	jc := permSpace[classNameStr]
	if jc == nil {
		jc = jclass.ParseInstanceByClassName(classNameStr)
	}
	utils.AssertTrue(jc != nil, "NullPointerException")
	methodInfo, err, _ := jc.FindMethod(methodNameStr, methodDescStr)
	utils.AssertTrue(err == nil, "no find the method of "+methodNameStr)
	utils.AssertFalse(jclass.IsStatic(methodInfo.AccessFlag()), "IncompatibleClassChangeError")

	if jclass.IsProteced(methodInfo.AccessFlag()) {
		// todo if is proteced , need to judge the relation between caller and called
	}

	base.InvokeMethod(frame, methodInfo, false)
}
