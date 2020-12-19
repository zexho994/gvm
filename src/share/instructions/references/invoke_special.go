package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用父类方法、实例初始化方法（<init>）、私有方法
type INVOKE_SPECIAL struct {
	base.InstructionIndex16
}

func (i INVOKE_SPECIAL) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	constantMethod := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantMethod)
	perm := jclass.GetPerm()
	jc := perm.Space[constantMethod.ClassName()]
	exception.AssertTrue(jc != nil, "Class uninitialized")
	name, Desc := constantMethod.NameAndDescriptor()
	method, _ := jc.FindMethod(name, Desc)
	// 如果是初始化方法
	base.InvokeMethod(frame, method)
}
