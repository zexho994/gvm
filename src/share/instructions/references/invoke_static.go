package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用静态方法
//
type INVOKE_STATIC struct {
	base.InstructionIndex16
}

func (i *INVOKE_STATIC) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	contantMethod := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantMethod)
	name, t := contantMethod.NameAndDescriptor()
	className := contantMethod.ClassName()

	fmt.Println(name, t, className)
}
