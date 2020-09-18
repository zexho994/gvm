package rtda

import (
	"./heap"
	"math"
)

type OperandStack struct {
	// record the top position of the stack
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint16) *OperandStack {
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

func (self *OperandStack) PushInt(val int32) {
	//fmt.Printf("[gvm][PushInt] 操作数栈push新值: val : %v \n", val)
	self.slots[self.size].num = val
	//fmt.Printf("[gvm][PushInt] 操作数栈push新值后的结果: val : %v \n", self.slots[self.size].num)
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	//fmt.Printf("[gvm][PushInt] 操作数栈pop,当前长度size: %v\n", self.size)
	self.size--
	//fmt.Printf("[gvm][PushInt] 操作数栈pop值：%v\n", self.slots[self.size].num)
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)

	return math.Float32frombits(bits)
}

func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)

	return int64(high)<<32 | int64(low)
}

func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())

	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}

/*
extend OperandStack size
the operandStack size + 1
*/
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

/*
reduce the OperandStack size
the operandStack size - 1
*/
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

/*
获取距离top n个距离的引用
比如GetRefFromTop(0)获取栈顶的引用
比如GetRefFromTop(1)获取距离栈顶1个单位长度的引用
*/
func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	targetIndex := self.size - 1 - n
	//fmt.Printf("[gvm][operand_stack.GetRefFromTop] stack size : %v , target index : %v \n", self.size, targetIndex)
	return self.slots[targetIndex].ref
}
