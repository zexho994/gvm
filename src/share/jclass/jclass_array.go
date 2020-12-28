package jclass

import "github.com/zouzhihao-994/gvm/src/share/exception"

const (
	T_BOOLEAN = 4
	T_CHAT    = 5
	T_FLOAT   = 6
	T_DOUBLE  = 7
	T_BYTE    = 8
	T_SHORT   = 9
	T_INT     = 10
	T_LONG    = 11
)

type arrayData interface {
	newArray(len uint32) interface{}
}

type iArray []int32
type lArray []int64

func (iarr iArray) newArray(len uint32) interface{} {
	return make([]int32, len)
}

func initArrayData(len uint32, atype uint8) (interface{}, *exception.GvmError) {
	switch atype {
	case 10:
		return iArray{}.newArray(len), nil
	default:
		return nil, &exception.GvmError{Msg: "Array format Error"}
	}
}

type JArray struct {
	length uint32
	atype  uint8
	data   arrayData
}

func NewJarray(len uint32, atype uint8) *JArray {
	arr, _ := initArrayData(len, atype)
	return &JArray{
		length: len,
		atype:  atype,
		data:   arr.(arrayData),
	}
}
