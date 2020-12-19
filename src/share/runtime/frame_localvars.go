package runtime

import (
	"github.com/zouzhihao-994/gvm/src/share/oops"
	"math"
)

type LocalVars struct {
	slots []Slot
}

func NewLocalVars(maxLocals uint16) *LocalVars {
	if maxLocals > 0 {
		localvars := &LocalVars{
			slots: make([]Slot, maxLocals),
		}
		return localvars
	}
	return nil
}

func (l *LocalVars) SetInt(index uint, val int32) {
	l.slots[index].num = val
}

func (l *LocalVars) GetInt(index uint) int32 {
	return l.slots[index].num
}

func (l *LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l.slots[index].num = int32(bits)
}

func (l *LocalVars) GetFloat(index uint) float32 {
	bits := uint32(l.slots[index].num)

	return math.Float32frombits(bits)
}

func (l *LocalVars) SetLong(index uint, val int64) {
	l.slots[index].num = int32(val)
	l.slots[index+1].num = int32(val >> 32)
}

func (l *LocalVars) GetLong(index uint) int64 {
	low := uint32(l.slots[index].num)
	high := uint32(l.slots[index+1].num)

	return int64(high)<<32 | int64(low)
}

func (l LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l LocalVars) GetDouble(index uint) float64 {
	bits := uint64(l.GetLong(index))

	return math.Float64frombits(bits)
}

func (l LocalVars) SetRef(index uint, ref *oops.Oop_Instance) {
	l.slots[index].ref = ref
}

func (l LocalVars) GetRef(index uint) *oops.Oop_Instance {
	return l.slots[index].ref
}

func (l LocalVars) SetSlot(index uint, slot Slot) {
	l.slots[index] = slot
}

func (l LocalVars) GetThis() *oops.Oop_Instance {
	return l.GetRef(0)
}
