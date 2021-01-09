package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/oops"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 创建一个reference类型的数组
// index指向当前类的的运行时常量池
// 对应项应该是一个类、接口、数组类型的符号引用，而且应该已经被解析
type ANEW_ARRAY struct {
	base.InstructionIndex16
}

func (i *ANEW_ARRAY) Execute(frame *runtime.Frame) {
	arrayLength := frame.OperandStack().PopInt()
	c := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantClassInfo)
	cname := c.Name()
	jc := jclass.GetPerm().Space[cname]
	if jc == nil {
		jc = jclass.ParseInstanceByClassName(cname)
	}
	jarry := oops.NewRefJarray(uint32(arrayLength), jc)
	arrayInstance := oops.NewArrayOopInstance(&jarry)
	frame.OperandStack().PushRef(arrayInstance)
}
