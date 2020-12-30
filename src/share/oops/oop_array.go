package oops

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
)

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
type cArray []int8

func (iarr iArray) newiArray(len uint32) iArray {
	return make([]int32, len)
}
func (carr cArray) newcArray(len uint32) cArray {
	return make([]int8, len)
}

type JArray struct {
	length uint32
	atype  uint8
	data   interface{} // array
}

//
func (jarray *JArray) SetIVal(idx int32, val int32) {
	exception.AssertTrue(jarray.atype == T_INT, "ArrayTypeError")
	exception.AssertTrue(idx >= 0 && idx < int32(jarray.length), "ArrayIndexOutBoundsException")
	ia := jarray.data.(iArray)
	ia[idx] = val
}
func (jarray *JArray) SetCVal(idx int32, val int8) {
	exception.AssertTrue(jarray.atype == T_CHAT, "ArrayTypeError")
	exception.AssertTrue(idx >= 0 && idx < int32(jarray.length), "ArrayIndexOutBoundsException")
	ia := jarray.data.(cArray)
	ia[idx] = val
}
func NewJarray(len uint32, atype uint8) *JArray {
	jarray := &JArray{
		length: len,
		atype:  atype,
	}
	switch atype {
	case 5:
		jarray.data = cArray{}.newcArray(len)
	case 10:
		jarray.data = iArray{}.newiArray(len)
	default:
		exception.GvmError{Msg: "NewJarrayError"}.Throw()
	}
	return jarray
}
