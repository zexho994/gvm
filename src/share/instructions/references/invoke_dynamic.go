package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用动态方法
type INVOKE_DYNAMIC struct {
	base.InstructionIndex32
}

// invokedynamic指令出现的地方称为 "动态调用点"
//
func (i *INVOKE_DYNAMIC) Execute(frame *runtime.Frame) {
	constantPool := frame.Method().CP()
	indexByte := uint16(i.Index >> 16)
	constInvokeDynamic := constantPool.GetConstantInfo(indexByte).(*constant_pool.ConstantInvokeDynamic)
	name, desc := constantPool.GetNameAndType(constInvokeDynamic.NameAndTypeIndex)
	fmt.Println(name, desc)

}
