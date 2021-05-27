package runtime

//import (
//	"github.com/zouzhihao-994/gvm/klass"
//	"github.com/zouzhihao-994/gvm/utils"
//)
//
//func NewShimFrame(thread *Thread, args []utils.Slot) *Frame {
//	return newShimFrame(thread, args)
//}
//
//func newShimFrame(thread *Thread, args []utils.Slot) *Frame {
//	return &Frame{
//		Thread:       thread,
//		MethodKlass:  klass.NewMethodKlass(utils.ACC_STATIC,),
//		OperandStack: newOperandStackWithSlots(args),
//	}
//}
//
//func newAthrowFrame(thread *Thread, ex *heap.Object, initArgs []heap.Slot) *Frame {
//	// stackSlots := [ex, ex, initArgs]
//	stackSlots := make([]heap.Slot, len(initArgs)+2)
//	stackSlots[0] = heap.NewRefSlot(ex)
//	stackSlots[1] = heap.NewRefSlot(ex)
//	copy(stackSlots[2:], initArgs)
//
//	return &Frame{
//		Thread:       thread,
//		Method:       shimAThrowMethod,
//		OperandStack: newOperandStackWithSlots(stackSlots),
//	}
//}
