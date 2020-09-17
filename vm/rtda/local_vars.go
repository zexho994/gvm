package rtda

import (
	"./heap"
	"fmt"
	"math"
)

type LocalVars struct {
	slots []Slot
}

func NewLocalVars(maxLocals uint16) *LocalVars {
	if maxLocals > 0 {
		fmt.Printf("[gvm][LocalVars.NewLocalVars] maxLocals : %v \n", maxLocals)
		localvars := &LocalVars{
			slots: make([]Slot, maxLocals),
		}
		fmt.Println("[gvm][NewLocalVars] done")
		return localvars
	}
	return nil
}

func (self *LocalVars) SetInt(index uint, val int32) {
	fmt.Printf("[gvm][SetInt] index : %v, val : %v \n", index, val)
	self.slots[index].num = val
	fmt.Printf("[gvm][SetInt] 设置后的结果 val : %v \n", self.GetInt(index))
}

func (self *LocalVars) GetInt(index uint) int32 {
	return self.slots[index].num
}

func (self *LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self.slots[index].num = int32(bits)
}

func (self *LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self.slots[index].num)

	return math.Float32frombits(bits)
}

func (self *LocalVars) SetLong(index uint, val int64) {
	self.slots[index].num = int32(val)
	self.slots[index+1].num = int32(val >> 32)
}

func (self *LocalVars) GetLong(index uint) int64 {
	low := uint32(self.slots[index].num)
	high := uint32(self.slots[index+1].num)

	return int64(high)<<32 | int64(low)
}

func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))

	return math.Float64frombits(bits)
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self.slots[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *heap.Object {
	return self.slots[index].ref
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
	self.slots[index] = slot
}
