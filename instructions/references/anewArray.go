package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
)

// AnewArray 创建一个reference类型的数组
// index指向当前类的的运行时常量池
// 对应项应该是一个类、接口、数组类型的符号引用，而且应该已经被解析
type AnewArray struct {
	base.InstructionIndex16
}

func (i *AnewArray) Execute(frame *runtime.Frame) {
	constantClassInfo := frame.Method().GetConstantInfo(i.Index).(*constant_pool.ConstantClassInfo)
	cname := constantClassInfo.Name()
	k := klass.Perm().Get(cname)
	if k == nil {
		k = klass.ParseByClassName(cname)
	}
	if !k.IsInit {
		frame.RevertPC()
		base.InitClass(k, frame.Thread)
		return
	}

	arrayLength := frame.PopInt()
	jarry := oops.NewRefJarray(uint32(arrayLength), k)
	arrayInstance := oops.NewArrayOopInstance(&jarry)
	frame.PushRef(arrayInstance)

}
