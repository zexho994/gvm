package heap

import (
	"math"
)

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (slots Slots) SetInt(index uint, val int32) {
	slots[index].num = val
}
func (slots Slots) GetInt(index uint) int32 {
	return slots[index].num
}

func (slots Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	slots[index].num = int32(bits)
}
func (slots Slots) GetFloat(index uint) float32 {
	bits := uint32(slots[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (slots Slots) SetLong(index uint, val int64) {
	slots[index].num = int32(val)
	slots[index+1].num = int32(val >> 32)
}
func (slots Slots) GetLong(index uint) int64 {
	low := uint32(slots[index].num)
	high := uint32(slots[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (slots Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	slots.SetLong(index, int64(bits))
}
func (slots Slots) GetDouble(index uint) float64 {
	bits := uint64(slots.GetLong(index))
	return math.Float64frombits(bits)
}

func (slots Slots) SetRef(index uint, ref *Object) {
	slots[index].ref = ref
}
func (slots Slots) GetRef(index uint) *Object {
	return slots[index].ref
}
