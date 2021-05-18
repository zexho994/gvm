package oops

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/utils"
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
type aArray []klass.Klass

func (iarr iArray) newiArray(len uint32) iArray {
	return make([]int32, len)
}

func (carr cArray) newcArray(len uint32) cArray {
	return make([]int8, len)
}

func (arr aArray) newaArray(len uint32) aArray {
	return make([]klass.Klass, len)
}

type JArray struct {
	length  uint32
	arrtype uint8
	refType interface{}
	data    interface{} // array
}

func (jarray *JArray) SetIVal(idx int32, val int32) {
	utils.AssertTrue(jarray.arrtype == T_INT, "ArrayTypeError")
	utils.AssertTrue(idx >= 0 && idx < int32(jarray.length), "ArrayIndexOutBoundsException")
	ia := jarray.data.(iArray)
	ia[idx] = val
}

func (jarray *JArray) SetCVal(idx int32, val int8) {
	utils.AssertTrue(jarray.arrtype == T_CHAT, "ArrayTypeError")
	utils.AssertTrue(idx >= 0 && idx < int32(jarray.length), "ArrayIndexOutBoundsException")
	ia := jarray.data.(cArray)
	ia[idx] = val
}

func NewRefJarray(len uint32, instance *klass.Klass) JArray {
	if len < 0 {
		exception.GvmError{Msg: "NegativeArraySizeException"}.Throw()
	}
	return JArray{
		length:  len,
		refType: instance,
	}
}

func NewJarray(len uint32, atype uint8) *JArray {
	jarray := &JArray{
		length:  len,
		arrtype: atype,
	}
	switch atype {
	case 5:
		jarray.data = cArray{}.newcArray(len)
	case 10:
		jarray.data = iArray{}.newiArray(len)
	}
	return jarray
}
