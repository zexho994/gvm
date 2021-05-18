package runtime

import (
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/utils"
	"math"
)

type OperandStack struct {
	// record the top position of the stack
	size  uint32
	slots utils.Slots
}

func NewOperandStack(maxStack uint16) *OperandStack {
	if maxStack > 0 {
		operandStack := &OperandStack{
			slots: make([]utils.Slot, maxStack),
		}
		return operandStack
	}

	return nil
}

func (operandStack *OperandStack) PushInt(val int32) {
	operandStack.slots[operandStack.size].Num = val
	operandStack.size++
}

func (operandStack *OperandStack) PopInt() int32 {
	operandStack.size--
	return operandStack.slots[operandStack.size].Num
}

func (operandStack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	operandStack.slots[operandStack.size].Num = int32(bits)
	operandStack.size++
}
func (operandStack *OperandStack) PopFloat() float32 {
	operandStack.size--
	bits := uint32(operandStack.slots[operandStack.size].Num)
	return math.Float32frombits(bits)
}

func (operandStack *OperandStack) PushLong(val int64) {
	operandStack.slots[operandStack.size].Num = int32(val)
	operandStack.slots[operandStack.size+1].Num = int32(val >> 32)
	operandStack.size += 2
}

func (operandStack *OperandStack) PopLong() int64 {
	operandStack.size -= 2
	low := uint32(operandStack.slots[operandStack.size].Num)
	high := uint32(operandStack.slots[operandStack.size+1].Num)
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

func (operandStack *OperandStack) PushRef(ref *oops.OopInstance) {
	operandStack.slots[operandStack.size].Ref = ref
	operandStack.size++
}

func (operandStack *OperandStack) PopRef() *oops.OopInstance {
	operandStack.size--
	ref := operandStack.slots[operandStack.size].Ref.(*oops.OopInstance)
	operandStack.slots[operandStack.size].Ref = nil
	return ref
}

// PushSlot /*
func (operandStack *OperandStack) PushSlot(slot utils.Slot) {
	operandStack.slots[operandStack.size] = slot
	operandStack.size++
}

// PopSlot /*
func (operandStack *OperandStack) PopSlot() utils.Slot {
	operandStack.size--
	return operandStack.slots[operandStack.size]
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

// PopByParamters todo: provide more parmes type
func (operandStack *OperandStack) PopByParamters(params []string, localVars *LocalVars, isStatic bool) {
	i := len(params)
	// method is storages <this.class.Ref> on localvars[0]
	// but the static_method is different,don't storages Ref on [0]
	if isStatic {
		i--
	}
	for idx := range params {
		switch params[idx] {
		case "B":
			break
		case "C":
			break
		case "D":
			localVars.SetDouble(uint(i-idx), operandStack.PopDouble())
			break
		case "F":
			localVars.SetFloat(uint(i-idx), operandStack.PopFloat())
			break
		case "I":
			localVars.SetInt(uint(i-idx), operandStack.PopInt())
			break
		case "J":
			operandStack.PopLong()
			localVars.SetLong(uint(i-idx), operandStack.PopLong())
			break
		case "S":
			break
		case "Z":
			operandStack.PopBoolean()
			localVars.SetBoolean(uint(i+idx), operandStack.PopBoolean())
			break
		case "L":
		case "[":
			operandStack.PopRef()
			break
		}
	}

	// save the invoke class Ref to localvars[0]
	if !isStatic {
		localVars.SetRef(0, operandStack.PopRef())
	}

}
