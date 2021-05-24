package utils

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/exception"
)

type Slots []Slot

type Slot struct {
	Num  int32
	Ref  interface{}
	Type uint8
}

const (
	SlotByte    = 1
	SlotChar    = 2
	SlotDouble  = 3
	SlotFloat   = 4
	SlotInt     = 5
	SlotLong    = 6
	SlotRef     = 7
	SlotShort   = 8
	SlotBoolean = 9
)

// GetVal64 获取64位值
// 对于操作数栈来说，一个64位数拆分成两个32位，并且高位先入栈
// return low，hight
func (slots Slots) GetVal64() (int32, int32) {
	return slots[0].Num, slots[1].Num
}

func TypeMapping(desc string) uint8 {
	switch desc {
	case config.Int:
		return SlotInt
	case config.Long:
		return SlotLong
	case config.Byte:
		return SlotByte
	case config.Double:
		return SlotDouble
	case config.Float:
		return SlotFloat
	case config.Char:
		return SlotChar
	case config.Boolean:
		return SlotBoolean
	default: // refrence type
		return SlotRef
	}
	exception.GvmError{Msg: "type mapping error,desc = " + desc}.Throw()
	return 0
}
