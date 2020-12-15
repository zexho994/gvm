package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用静态方法
//
type INVOKE_STATIC struct {
	base.InstructionIndex16
}

// invoke a static method
// get the static method -> verify the access flag of method equals static
// -> parse constant method to
func (i *INVOKE_STATIC) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	contantMethod := cp.GetConstantInfo(i.Index).(*constant_pool.ConstantMethod)
	className := contantMethod.ClassName()
	perm := jclass.GetPerm()
	class := perm.Space[className]
	if class == nil {
		class = jclass.ParseInstanceByClassName(className)
	}

}
