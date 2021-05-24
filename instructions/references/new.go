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

// NEW 创建一个实例
type NEW struct {
	base.InstructionIndex16
}

func (n *NEW) Execute(frame *runtime.Frame) {
	// 获取类常量信息
	constantClass := frame.GetConstantInfo(n.Index).(*constant_pool.ConstantClassInfo)
	className := constantClass.Name()

	// 判断类是否已经加载过
	class := klass.Perm.Get(className)
	if class == nil {
		class = klass.ParseByClassName(className)
		klass.Perm.Save(className, class)
	}
	if !class.IsInit {
		base.InitClass(class, frame.Thread)
	}
	if utils.IsInterface(class.AccessFlags) || utils.IsAbstract(class.AccessFlags) {
		panic(exception.GvmError{Msg: "[gvm] the interface and abstract cannot be instantiated"})
	}

	instance := oops.NewOopInstance(class)
	frame.PushRef(instance)

}
