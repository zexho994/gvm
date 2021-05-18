package references

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/heap"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
)

// 创建一个实例
type NEW struct {
	base.InstructionIndex16
}

func (n *NEW) Execute(frame *runtime.Frame) {
	// 获取类常量信息
	cp := frame.Method().CP()
	constantClass := cp.GetConstantInfo(n.Index).(*constant_pool.ConstantClassInfo)
	className := constantClass.Name()

	// 判断类是否已经加载过
	perm := jclass.Perm()
	class := perm.Space[className]

	// 还未加载过
	if class == nil {
		class = jclass.ParseInstanceByClassName(className)
		perm.Space[className] = class
	}

	// 判断类是否初始化过
	if !class.IsInit {
		base.InitClass(class, frame.Thread())
	}

	// 初始化一个实例
	if jclass.IsInterface(class.AccessFlags) || jclass.IsAbstract(class.AccessFlags) {
		panic(exception.GvmError{Msg: "[gvm] the interface and abstract cannot be instantiated"})
	}
	instance := oops.NewOopInstance(class)
	heap.GetHeap().Add(instance)
	frame.OperandStack().PushRef(instance)

}
