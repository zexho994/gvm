package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

type InvokeInterface struct {
	base.InstructionIndex32
}

// Execute indexbyte1,indexbyte2,count,0
// indexbyte1 indexbyte2指向常量的索引
// count无符号类型，非0
// 该方法不能是实例初始化方法<init>、类或接口初始化方法<clinit>
func (i InvokeInterface) Execute(frame *runtime.Frame) {
	poolIndex := i.Index >> 16
	count := (i.Index << 16) >> 16
	fmt.Println(count)
	constantMethod := frame.Method().CP().GetConstantInfo(uint16(poolIndex)).(*constant_pool.ConstantInterfaceMethodInfo)
	methodNameStr, methodDescStr := constantMethod.NameAndDescriptor()
	k := frame.OperandStack().GetByIdx(0)
	methodInfo, err, _ := k.Ref.(*oops.OopInstance).Klass().FindMethod(methodNameStr, methodDescStr)
	utils.AssertError(err, "no find the method of "+methodNameStr)

	base.InvokeMethod(frame, methodInfo, false)
}
