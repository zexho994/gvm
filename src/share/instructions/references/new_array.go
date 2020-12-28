package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

const (
	T_BOOLEAN = 4
	T_CHAT    = 5
	T_FLOAT   = 6
	T_DOUBLE  = 7
	T_BYTE    = 8
	T_SHORT   = 9
	T_INT     = 10
	T_LONG    = 11
)

type NEW_ARRAY struct {
	count    uint
	arrayref *jclass.JClass_Instance
	base.InstructionIndex8
}

func (i *NEW_ARRAY) Execute(frame *runtime.Frame) {
	arrayCount := frame.OperandStack().PopInt()
	exception.AssertFalse(arrayCount < 0, "NegativeArraySizeException")

}
