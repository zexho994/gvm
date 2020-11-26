package runtime

import (
	"github.com/zouzhihao-994/gvm/src/vm/oops"
	"math"
)

type OperandStack struct {
	// record the top position of the stack
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	//fmt.Printf("[gvm][OperandStack.NewOperandStack] maxStack : %v \n", maxStack)
	if maxStack > 0 {
		operandStack := &OperandStack{
			slots: make([]Slot, maxStack),
		}
		//fmt.Printf("[gvm][OperandStack.NewOperandStack] done \n")
		return operandStack
	}

	return nil
}

func (operandStack *OperandStack) PushInt(val int32) {
	//fmt.Printf("[gvm][PushInt] 操作数栈push新值: val : %v \n", val)
	operandStack.slots[operandStack.size].num = val
	//fmt.Printf("[gvm][PushInt] 操作数栈push新值后的结果: val : %v \n", operandStack.slots[operandStack.size].num)
	operandStack.size++
}

func (operandStack *OperandStack) PopInt() int32 {
	//fmt.Printf("[gvm][PushInt] 操作数栈pop,当前长度size: %v\n", operandStack.size)
	operandStack.size--
	//fmt.Printf("[gvm][PushInt] 操作数栈pop值：%v\n", operandStack.slots[operandStack.size].num)
	return operandStack.slots[operandStack.size].num
}

func (operandStack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	operandStack.slots[operandStack.size].num = int32(bits)
	operandStack.size++
}
func (operandStack *OperandStack) PopFloat() float32 {
	operandStack.size--
	bits := uint32(operandStack.slots[operandStack.size].num)

	return math.Float32frombits(bits)
}

func (operandStack *OperandStack) PushLong(val int64) {
	operandStack.slots[operandStack.size].num = int32(val)
	operandStack.slots[operandStack.size+1].num = int32(val >> 32)
	operandStack.size += 2
}

func (operandStack *OperandStack) PopLong() int64 {
	operandStack.size -= 2
	low := uint32(operandStack.slots[operandStack.size].num)
	high := uint32(operandStack.slots[operandStack.size+1].num)

	return int64(high)<<32 | int64(low)
}

func (operandStack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	operandStack.PushLong(int64(bits))
}

func (operandStack *OperandStack) PopDouble() float64 {
	bits := uint64(operandStack.PopLong())

	return math.Float64frombits(bits)
}

func (operandStack *OperandStack) PushRef(ref *oops.Object) {
	operandStack.slots[operandStack.size].ref = ref
	operandStack.size++
}

func (operandStack *OperandStack) PopRef() *oops.Object {
	operandStack.size--
	ref := operandStack.slots[operandStack.size].ref
	operandStack.slots[operandStack.size].ref = nil
	return ref
}

/*
extend OperandStack size
the operandStack size + 1
*/
func (operandStack *OperandStack) PushSlot(slot Slot) {
	operandStack.slots[operandStack.size] = slot
	operandStack.size++
}

/*
reduce the OperandStack size
the operandStack size - 1
*/
func (operandStack *OperandStack) PopSlot() Slot {
	operandStack.size--
	return operandStack.slots[operandStack.size]
}

/*
获取距离top n个距离的引用
比如GetRefFromTop(0)获取栈顶的引用
比如GetRefFromTop(1)获取距离栈顶1个单位长度的引用
*/
func (operandStack *OperandStack) GetRefFromTop(n uint) *oops.Object {
	targetIndex := operandStack.size - 1 - n
	//fmt.Printf("[gvm][operand_stack.GetRefFromTop] stack size : %v , target index : %v \n", operandStack.size, targetIndex)
	return operandStack.slots[targetIndex].ref
}

func (operandStack *OperandStack) PushBoolean(val bool) {
	if val {
		operandStack.PushInt(1)
	} else {
		operandStack.PushInt(0)
	}
}

func (operandStack *OperandStack) PopBoolean() bool {
	return operandStack.PopInt() == 1
}
