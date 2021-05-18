package references

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
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
	perm := klass.Perm()
	class := perm.Space[className]

	// 还未加载过
	if class == nil {
		class = klass.ParseByClassName(className)
		perm.Space[className] = class
	}

	// 判断类是否初始化过
	if !class.IsInit {
		base.InitClass(class, frame.Thread())
	}

	// 初始化一个实例
	if utils.IsInterface(class.AccessFlags) || utils.IsAbstract(class.AccessFlags) {
		panic(exception.GvmError{Msg: "[gvm] the interface and abstract cannot be instantiated"})
	}
	instance := oops.NewOopInstance(class)
	//heap.GetHeap().Add(instance)
	frame.OperandStack().PushRef(instance)

}
