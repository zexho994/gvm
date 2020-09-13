package heap

import (
	"fmt"
	"math"
)

type Slot struct {
	num int32
	ref *Object
}

type Slots struct {
	slots []Slot
}

func newSlots(slotCount uint) Slots {
	return Slots{slots: make([]Slot, slotCount)}
}

func (self Slots) SetInt(index uint, val int32) {
	fmt.Printf("[gvm][SetInt] index : %v, val : %v \n", index, val)
	self.slots[index].num = val
	fmt.Printf("[gvm][SetInt] 设置后的结果 val : %v \n", self.GetInt(index))
}

func (self Slots) GetInt(index uint) int32 {
	return self.slots[index].num
}

func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self.slots[index].num = int32(bits)
}

func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self.slots[index].num)

	return math.Float32frombits(bits)
}

func (self Slots) SetLong(index uint, val int64) {
	self.slots[index].num = int32(val)
	self.slots[index+1].num = int32(val >> 32)
}

func (self Slots) GetLong(index uint) int64 {
	low := uint32(self.slots[index].num)
	high := uint32(self.slots[index+1].num)

	return int64(high)<<32 | int64(low)
}

func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))

	return math.Float64frombits(bits)
}

func (self Slots) SetRef(index uint, ref *Object) {
	self.slots[index].ref = ref
}

func (self Slots) GetRef(index uint) *Object {
	return self.slots[index].ref
}
