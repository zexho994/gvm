package runtime

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"math"
)

type LocalVars struct {
	slots []Slot
}

func NewLocalVars(maxLocals uint) *LocalVars {
	if maxLocals > 0 {
		localvars := &LocalVars{
			slots: make([]Slot, maxLocals),
		}
		return localvars
	}
	return nil
}

func (localVars *LocalVars) SetInt(index uint, val int32) {
	localVars.slots[index].num = val
}

func (localVars *LocalVars) GetInt(index uint) int32 {
	return localVars.slots[index].num
}

func (localVars *LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	localVars.slots[index].num = int32(bits)
}

func (localVars *LocalVars) GetFloat(index uint) float32 {
	bits := uint32(localVars.slots[index].num)

	return math.Float32frombits(bits)
}

func (localVars *LocalVars) SetLong(index uint, val int64) {
	localVars.slots[index].num = int32(val)
	localVars.slots[index+1].num = int32(val >> 32)
}

func (localVars *LocalVars) GetLong(index uint) int64 {
	low := uint32(localVars.slots[index].num)
	high := uint32(localVars.slots[index+1].num)

	return int64(high)<<32 | int64(low)
}

func (localVars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	localVars.SetLong(index, int64(bits))
}

func (localVars LocalVars) GetDouble(index uint) float64 {
	bits := uint64(localVars.GetLong(index))

	return math.Float64frombits(bits)
}

func (localVars LocalVars) SetRef(index uint, ref *jclass.JClass_Instance) {
	localVars.slots[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint) *jclass.JClass_Instance {
	return localVars.slots[index].ref
}

func (localVars LocalVars) SetSlot(index uint, slot Slot) {
	localVars.slots[index] = slot
}

func (localVars LocalVars) GetThis() *jclass.JClass_Instance {
	return localVars.GetRef(0)
}
