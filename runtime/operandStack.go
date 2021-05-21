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

func (stack *OperandStack) GetByIdx(i int) utils.Slot {
	return stack.slots[i]
}

func (stack *OperandStack) SetByIdx(i int, s utils.Slot) {
	stack.slots[i] = s
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].Num = val
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	stack.size--
	return stack.slots[stack.size].Num
}

func (stack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	stack.slots[stack.size].Num = int32(bits)
	stack.size++
}
func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	bits := uint32(stack.slots[stack.size].Num)
	return math.Float32frombits(bits)
}

func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size].Num = int32(val)
	stack.slots[stack.size+1].Num = int32(val >> 32)
	stack.size += 2
}

func (stack *OperandStack) PopLong() int64 {
	stack.size -= 2
	low := uint32(stack.slots[stack.size].Num)
	high := uint32(stack.slots[stack.size+1].Num)
	return int64(high)<<32 | int64(low)
}

func (stack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	stack.PushLong(int64(bits))
}

func (stack *OperandStack) PopDouble() float64 {
	bits := uint64(stack.PopLong())
	return math.Float64frombits(bits)
}

func (stack *OperandStack) PushRef(ref *oops.OopInstance) {
	stack.slots[stack.size].Ref = ref
	stack.size++
}

func (stack *OperandStack) PopRef() *oops.OopInstance {
	stack.size--
	ref := stack.slots[stack.size].Ref.(*oops.OopInstance)
	stack.slots[stack.size].Ref = nil
	return ref
}

// PushSlot /*
func (stack *OperandStack) PushSlot(slot utils.Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}

// PopSlot /*
func (stack *OperandStack) PopSlot() utils.Slot {
	stack.size--
	return stack.slots[stack.size]
}

func (stack *OperandStack) PushBoolean(val bool) {
	if val {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

func (stack *OperandStack) PopBoolean() bool {
	return stack.PopInt() == 1
}

// PopByParamters todo: provide more parmes type
func (stack *OperandStack) PopByParamters(params []string, localVars *LocalVars, isStatic bool) {
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
			localVars.SetDouble(uint(i-idx), stack.PopDouble())
			break
		case "F":
			localVars.SetFloat(uint(i-idx), stack.PopFloat())
			break
		case "I":
			localVars.SetInt(uint(i-idx), stack.PopInt())
			break
		case "J":
			stack.PopLong()
			localVars.SetLong(uint(i-idx), stack.PopLong())
			break
		case "S":
			break
		case "Z":
			stack.PopBoolean()
			localVars.SetBoolean(uint(i+idx), stack.PopBoolean())
			break
		case "L":
		case "[":
			stack.PopRef()
			break
		}
	}

	// save the invoke class Ref to localvars[0]
	if !isStatic {
		localVars.SetRef(0, stack.PopRef())
	}

}
