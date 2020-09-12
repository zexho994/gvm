package heap

import (
	"fmt"
	"math"
)

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

type LocalVars struct {
	slots []Slot
}

func newSlots(slotCount uint) Slots {
	return make([]Slot, slotCount)
}

func NewLocalVars(maxLocals uint) *LocalVars {
	if maxLocals > 0 {
		fmt.Printf("[gvm][NewLocalVars] maxLocals : %v \n", maxLocals)
		return &LocalVars{
			slots: make([]Slot, maxLocals),
		}
	}
	return nil
}

func (self Slots) SetInt(index uint, val int32) {
	fmt.Printf("[gvm][SetInt] index : %v, val : %v \n", index, val)
	self[index].num = val
	fmt.Printf("[gvm][SetInt] 设置后的结果 val : %v \n", self.GetInt(index))
}

func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)

	return math.Float32frombits(bits)
}

func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)

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
	self[index].ref = ref
}

func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
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

func (self LocalVars) SetRef(index uint, ref *Object) {
	self.slots[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *Object {
	return self.slots[index].ref
}
