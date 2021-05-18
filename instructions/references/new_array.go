package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

type NEW_ARRAY struct {
	count    uint
	arrayref *klass.Klass
	base.InstructionIndex8
}

func (i *NEW_ARRAY) Execute(frame *runtime.Frame) {
	arrayCount := frame.OperandStack().PopInt()
	utils.AssertFalse(arrayCount < 0, "NegativeArraySizeException")
	arrayData := oops.NewJarray(uint32(arrayCount), i.Index)
	arrayOop := oops.NewArrayOopInstance(arrayData)
	//heap.GetHeap().Add(arrayOop)
	frame.OperandStack().PushRef(arrayOop)
}
