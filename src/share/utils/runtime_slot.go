package utils

import "github.com/zouzhihao-994/gvm/src/share/exception"

type Slots []Slot

type SlotVars struct {
	slots []Slot
}

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

// 对于一个64长度值
// d1 = 低32位
// d2 = 高32位
func (slots Slots) SetVal64(d1, d2 int32) {
	slots[0].Num = d1
	slots[1].Num = d2
}

func TypeMapping(desc string) uint8 {
	switch desc {
	case "I":
		return Slot_Int
	case "L":
		return Slot_Long
	case "B":
		return Slot_Byte
	case "D":
		return Slot_Double
	case "F":
		return Slot_Float
	case "J":
		return Slot_Long
	case "S":
		return Slot_Char
	case "Z":
		return Slot_Boolean
	default: // refrence type
		return Slot_Ref
	}
	exception.GvmError{Msg: "type mapping error,desc = " + desc}.Throw()
	return 0
}
