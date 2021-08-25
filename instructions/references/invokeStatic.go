package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	klass "github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InvokeStatic 调用静态方法
type InvokeStatic struct {
	base.InstructionIndex16
}

// Execute invoke a static method
// get the static method -> verify the access flag of method equals static -> parse constant method to
func (i *InvokeStatic) Execute(frame *runtime.Frame) {
	k := frame.GetConstantInfo(i.Index)
	var method *klass.MethodKlass
	var kl *klass.Klass
	if kMethodRef, ok := k.(*constant_pool.ConstantMethodInfo); ok {
		kMethodRef = k.(*constant_pool.ConstantMethodInfo)
		kl = klass.Perm.Get(kMethodRef.ClassName())
		if kl == nil {
			kl = klass.ParseByClassName(kMethodRef.ClassName())
		}
		if !kl.IsInit {
			frame.RevertPC()
			base.InitClass(kl, frame.Thread)
			return
		}
		method, _ = kl.FindStaticMethod(kMethodRef.NameAndDescriptor())
	} else {
		kMethodRef := k.(*constant_pool.ConstantInterfaceMethodInfo)
		kl = klass.Perm.Get(kMethodRef.ClassName())
		if kl == nil {
			kl = klass.ParseByClassName(kMethodRef.ClassName())
			klass.Perm.Save(kMethodRef.ClassName(), kl)
		}
		if !kl.IsInit {
			frame.RevertPC()
			base.InitClass(kl, frame.Thread)
			return
		}
		method, _ = kl.FindStaticMethod(kMethodRef.NameAndDescriptor())
	}

	base.InvokeMethod(frame, method, true)
}
