package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass/attribute"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 调用动态方法
type InvokeDynamic struct {
	base.InstructionIndex32
}

// invokedynamic指令出现的地方称为 "动态调用点"
// 解析出引导方法
func (i *InvokeDynamic) Execute(frame *runtime.Frame) {
	constantPool := frame.Method().CP()
	indexByte := uint16(i.Index >> 16)
	constInvokeDynamic := constantPool.GetConstantInfo(indexByte).(*constant_pool.ConstantInvokeDynamic)
	btmIdx := constInvokeDynamic.BootstrapMethodAttrIndex
	name, desc := constantPool.GetNameAndType(constInvokeDynamic.NameAndTypeIndex)
	fmt.Println(name, desc)
	attributeInfo, _ := frame.Method().JClass().Attributes.FindAttrInfo("BootstrapMethods")
	bsma := attributeInfo.(*attribute.BootstrapmethodsAttribute)
	// bootstrap_method
	bsm := bsma.Methods()[btmIdx]
	// bootstrap_method_ref
	methodHandle := constantPool.GetConstantInfo(bsm.MethodRef).(*constant_pool.ConstantMethodHandleInfo)
	methodHandle.ParseKindRef()

	fmt.Println(methodHandle)
}