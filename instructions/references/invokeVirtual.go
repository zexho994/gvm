package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// InvokeVirtual 调用实例方法
type InvokeVirtual struct {
	base.InstructionIndex16
}

func (i *InvokeVirtual) Execute(frame *runtime.Frame) {
	constantMethod := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantMethodInfo)
	methodNameStr, methodDescStr := constantMethod.NameAndDescriptor()
	utils.AssertTrue(methodNameStr != "<init>" && methodNameStr != "<clinit>", "IncompatibleClassChangeError")

	classNameStr := constantMethod.ClassName()
	k := klass.Perm().Space[classNameStr]
	if k == nil {
		k = klass.ParseByClassName(classNameStr)
	}
	utils.AssertTrue(k != nil, "NullPointerException")
	methodInfo, err, _ := k.FindMethod(methodNameStr, methodDescStr)
	utils.AssertTrue(err == nil, "no find the method of "+methodNameStr)
	utils.AssertFalse(utils.IsStatic(methodInfo.AccessFlag()), "IncompatibleClassChangeError")

	if utils.IsProteced(methodInfo.AccessFlag()) {
		// todo if is proteced , need to judge the relation between caller and called
	}

	base.InvokeMethod(frame, methodInfo, false)
}
