package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// InvokeSpecial 调用父类方法、实例初始化方法（<init>）、私有方法
type InvokeSpecial struct {
	base.InstructionIndex16
}

func (i *InvokeSpecial) Execute(frame *runtime.Frame) {
	if config.LogInvoke {
		fmt.Printf("----%s.%s%s class exec -> invokeSpecial ----\n",
			frame.ThisClass, frame.MethodName(), frame.MethodDescriptor())
	}
	cp := frame.ConstantPool
	k := cp.GetConstantInfo(i.Index)
	var kl *klass.Klass
	var method *klass.MethodKlass

	if kMethodRef, ok := k.(*constant_pool.ConstantMethodInfo); ok {
		kl = klass.Perm.Get(kMethodRef.ClassName())
		method, _, _ = kl.FindMethod(kMethodRef.NameAndDescriptor())
	} else {
		kMethodRef := k.(*constant_pool.ConstantInterfaceMethodInfo)
		kl = klass.Perm.Get(kMethodRef.ClassName())
		method, _, _ = kl.FindMethod(kMethodRef.NameAndDescriptor())
	}

	base.InvokeMethod(frame, method, utils.IsStatic(method.AccessFlag()))
}
