package references

import (
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
	cp := frame.Method().CP()
	constantMethod := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantMethodInfo)
	k := klass.Perm().Space[constantMethod.ClassName()]

	utils.AssertTrue(k != nil, "Class uninitialized")
	name, Desc := constantMethod.NameAndDescriptor()
	method, _, _ := k.FindMethod(name, Desc)

	// 如果是初始化方法
	base.InvokeMethod(frame, method, utils.IsStatic(method.AccessFlag()))
}
