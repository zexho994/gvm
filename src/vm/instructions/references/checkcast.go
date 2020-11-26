package references

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"
import "github.com/zouzhihao-994/gvm/src/vm/oops" // Check whether object is of given type

type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()

	classRef := cp.GetConstant(self.Index).(*oops.ClassRef)

	class := classRef.ResolvedClass()

	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}

}
