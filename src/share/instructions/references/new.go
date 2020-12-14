package references

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 创建一个实例
type NEW struct {
	base.InstructionIndex16
}

func (n NEW) Execute(frame *runtime.Frame) {
	cp := frame.Method().CP()
	// 获取类常量信息
	constantClass := cp.GetConstantInfo(n.Index).(*constant_pool.ConstantClass)
	className := constantClass.Name()
	// 判断类是否已经加载过
	perm := jclass.GetPerm()
	name := perm.Space[className]
	// 还未加载过
	if name == nil {

	}

}
