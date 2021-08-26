package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// InvokeVirtual 调用实例方法
type InvokeVirtual struct {
	base.InstructionIndex16
}

func (i *InvokeVirtual) Execute(frame *runtime.Frame) {
	if config.LogInvoke {
		fmt.Printf("----%s.%s%s class exec -> invokevirtual ----\n",
			frame.ThisClass, frame.MethodName(), frame.MethodDescriptor())
	}
	constantMethod := frame.GetConstantMethodInfo(i.Index)
	objectRef := frame.GetByIdx(0)
	k := objectRef.Ref.(*oops.OopInstance).Klass
	method, err, _ := k.FindMethod(constantMethod.NameAndDescriptor())
	utils.AssertError(err, "klass to find method err")

	base.InvokeMethod(frame, method, false)
}
