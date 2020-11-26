package references

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"
import "github.com/zouzhihao-994/gvm/src/vm/oops"

// Create new array of reference
// 引用类型数组需要两个操作数
// 第一个16位操作数用来表示从当前类的运行时常量池中找到类符号引用
// 第二个操作数是数组长度
type ANEW_ARRAY struct{ base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 根据第一个操作数在常量池中获取类的符号引用
	classRef := cp.GetConstant(self.Index).(*oops.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
