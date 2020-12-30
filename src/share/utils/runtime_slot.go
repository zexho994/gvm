package utils

type Slots []Slot

type Slot struct {
	Num  int32
	Ref  interface{}
	Type uint8
}

const (
	Slot_Byte    = 1
	Slot_Char    = 2
	Slot_Double  = 3
	Slot_Float   = 4
	Slot_Int     = 5
	Slot_Long    = 6
	Slot_Ref     = 7
	Slot_Short   = 8
	Slot_Boolean = 9
)

func (slots Slots) SetSlot(idx uint32, s Slot) {
	slots[idx] = s
}
func (slots Slots) SetInt(idx uint32, val int32) {
	slots[idx].Num = val
	slots[idx].Type = Slot_Int
}
