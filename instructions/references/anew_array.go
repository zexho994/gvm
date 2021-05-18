package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
)

// ANEW_ARRAY 创建一个reference类型的数组
// index指向当前类的的运行时常量池
// 对应项应该是一个类、接口、数组类型的符号引用，而且应该已经被解析
type ANEW_ARRAY struct {
	base.InstructionIndex16
}

func (i *ANEW_ARRAY) Execute(frame *runtime.Frame) {
	arrayLength := frame.OperandStack().PopInt()
	c := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantClassInfo)
	cname := c.Name()
	jc := klass.Perm().Space[cname]
	if jc == nil {
		jc = klass.ParseInstanceByClassName(cname)
	}
	jarry := oops.NewRefJarray(uint32(arrayLength), jc)
	arrayInstance := oops.NewArrayOopInstance(&jarry)
	frame.OperandStack().PushRef(arrayInstance)
}
