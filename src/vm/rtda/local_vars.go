package rtda

import (
	"github.com/zouzhihao-994/gvm/src/vm/rtda/heap"
	"math"
)

type LocalVars struct {
	slots []Slot
}

func NewLocalVars(maxLocals uint) *LocalVars {
	if maxLocals > 0 {
		//fmt.Printf("[gvm][LocalVars.NewLocalVars] maxLocals : %v \n", maxLocals)
		localvars := &LocalVars{
			slots: make([]Slot, maxLocals),
		}
		//fmt.Println("[gvm][NewLocalVars] done")
		return localvars
	}
	return nil
}

func (localVars *LocalVars) SetInt(index uint, val int32) {
	//fmt.Printf("[gvm][SetInt] index : %v, val : %v \n", index, val)
	localVars.slots[index].num = val
	//fmt.Printf("[gvm][SetInt] 设置后的结果 val : %v \n", localVars.GetInt(index))
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

func (localVars LocalVars) SetRef(index uint, ref *heap.Object) {
	localVars.slots[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint) *heap.Object {
	return localVars.slots[index].ref
}

func (localVars LocalVars) SetSlot(index uint, slot Slot) {
	localVars.slots[index] = slot
}

func (localVars LocalVars) GetThis() *heap.Object {
	return localVars.GetRef(0)
}
