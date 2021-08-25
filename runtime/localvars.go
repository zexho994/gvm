package runtime

import (
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/utils"
	"math"
)

type LocalVars struct {
	slots []utils.Slot
}

func NewLocalVars(maxLocals uint16) *LocalVars {
	if maxLocals > 0 {
		localvars := &LocalVars{
			slots: make([]utils.Slot, maxLocals),
		}
		return localvars
	}
	return nil
}

func (l *LocalVars) SetInt(index uint, val int32) {
	l.slots[index].Num = val
	l.slots[index].Type = utils.SlotInt
}

func (l *LocalVars) GetInt(index uint) int32 {
	return l.slots[index].Num
}

func (l *LocalVars) SetBoolean(index uint, val bool) {
	var n int32
	if val {
		n = 1
	} else {
		n = 0
	}
	l.slots[index].Num = n
	l.slots[index].Type = utils.SlotBoolean
}

func (l *LocalVars) GetBoolean(index uint) bool {
	return l.slots[index].Num == 1
}

func (l *LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l.slots[index].Num = int32(bits)
	l.slots[index].Type = utils.SlotFloat

}

func (l *LocalVars) GetFloat(index uint) float32 {
	bits := uint32(l.slots[index].Num)

	return math.Float32frombits(bits)
}

func (l *LocalVars) SetLong(index uint, val int64) {
	l.slots[index].Num = int32(val)
	l.slots[index+1].Num = int32(val >> 32)
	l.slots[index].Type = utils.SlotLong

}

func (l *LocalVars) GetLong(index uint) int64 {
	low := uint32(l.slots[index].Num)
	high := uint32(l.slots[index+1].Num)

	return int64(high)<<32 | int64(low)
}

func (l *LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
	l.slots[index].Type = utils.SlotDouble
}

func (l *LocalVars) GetDouble(index uint) float64 {
	bits := uint64(l.GetLong(index))
	return math.Float64frombits(bits)
}

func (l *LocalVars) SetRef(index uint, ref *oops.OopInstance) {
	l.slots[index].Ref = ref
	l.slots[index].Type = utils.SlotRef

}

func (l *LocalVars) GetRef(index uint) *oops.OopInstance {
	return l.slots[index].Ref.(*oops.OopInstance)
}

func (l *LocalVars) SetSlot(index uint, slot utils.Slot) {
	l.slots[index] = slot
}

func (l *LocalVars) GetThis() *oops.OopInstance {
	return l.GetRef(0)
}
