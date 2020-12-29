package oops

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

type iArray []int32

func (iarr iArray) newArray(len uint32) interface{} {
	return make([]int32, len)
}

type JArray struct {
	length uint32
	atype  uint8
	data   interface{} // array
}

func NewJarray(len uint32, atype uint8) *JArray {
	jarray := &JArray{
		length: len,
		atype:  atype,
	}
	switch atype {
	case 10:
		jarray.data = iArray{}.newArray(len)
	default:
		exception.GvmError{Msg: "NewJarrayError"}.Throw()
	}
	return jarray
}
