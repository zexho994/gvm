package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InvokeDynamic 调用动态方法
type InvokeDynamic struct {
	base.InstructionIndex32
}

// Execute invokedynamic指令出现的地方称为 "动态调用点"
// 解析出引导方法
func (i *InvokeDynamic) Execute(frame *runtime.Frame) {
	constantPool := frame.ConstantPool
	indexByte := uint16(i.Index >> 16)
	constInvokeDynamic := constantPool.GetConstantDynamicInfo(indexByte)
	btmIdx := constInvokeDynamic.BootstrapMethodAttrIndex
	name, desc := constantPool.GetNameAndType(constInvokeDynamic.NameAndTypeIndex)
	fmt.Println(name, desc)
	attributeInfo, _ := frame.Klass.AttributesInfo.FindAttrInfo("BootstrapMethods")
	bsma := attributeInfo.(*attribute.BootstrapmethodsAttribute)
	// bootstrap_method
	bsm := bsma.Methods()[btmIdx]
	// bootstrap_method_ref
	methodHandle := constantPool.GetConstantMethodHandleInfo(bsm.MethodRef)
	methodHandle.ParseKindRef()

	fmt.Println(methodHandle)
}
