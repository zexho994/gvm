package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InvokeDynamic 调用动态方法
type InvokeDynamic struct {
	base.InstructionIndex32
}

// Execute invokedynamic指令出现的地方称为 "动态调用点"
// 解析出引导方法
func (i *InvokeDynamic) Execute(frame *runtime.Frame) {
	constantPool := frame.Method().CP()
	indexByte := uint16(i.Index >> 16)
	constInvokeDynamic := constantPool.GetConstantInfo(indexByte).(*constant_pool.ConstantInvokeDynamic)
	btmIdx := constInvokeDynamic.BootstrapMethodAttrIndex
	name, desc := constantPool.GetNameAndType(constInvokeDynamic.NameAndTypeIndex)
	fmt.Println(name, desc)
	attributeInfo, _ := frame.Method().Klass().Attributes.FindAttrInfo("BootstrapMethods")
	bsma := attributeInfo.(*attribute.BootstrapmethodsAttribute)
	// bootstrap_method
	bsm := bsma.Methods()[btmIdx]
	// bootstrap_method_ref
	methodHandle := constantPool.GetConstantInfo(bsm.MethodRef).(*constant_pool.ConstantMethodHandleInfo)
	methodHandle.ParseKindRef()

	fmt.Println(methodHandle)
}
