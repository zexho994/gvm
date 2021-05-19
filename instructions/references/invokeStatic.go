package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// INVOKE_STATIC 调用静态方法
type INVOKE_STATIC struct {
	base.InstructionIndex16
}

// Execute invoke a static method
// get the static method -> verify the access flag of method equals static
// -> parse constant method to
func (i *INVOKE_STATIC) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	contantMethod := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantMethodInfo)
	className := contantMethod.ClassName()
	perm := klass.Perm()
	class := perm.Space[className]
	if class == nil {
		class = klass.ParseByClassName(className)
	}
	name, _type := contantMethod.NameAndDescriptor()
	methodInfo, err := class.FindStaticMethod(name, _type)
	if err != nil {
		panic("[gvm]" + err.Error())
	}
	if !utils.IsStatic(methodInfo.AccessFlag()) {
		panic("[gvm] invoke static error")
	}
	base.InvokeMethod(frame, methodInfo, true)
}
