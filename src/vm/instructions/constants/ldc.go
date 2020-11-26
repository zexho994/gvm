package constants

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"
import "github.com/zouzhihao-994/gvm/src/vm/oops"

/*
ldc instruction is loading a varible from constants_pool and pushed to openstack
*/
type LDC struct{ base.Index8Instruction }

/**

 */
type LDC_W struct{ base.Index16Instruction }

type LDC2_W struct{ base.Index16Instruction }

func (self *LDC) Execute(frame *runtime.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *runtime.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *runtime.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := oops.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *oops.ClassRef:
		classRef := c.(*oops.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

func (self *LDC2_W) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
