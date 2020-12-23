package runtime

import (
	"github.com/zouzhihao-994/gvm/src/share/oops"
)

type Slot struct {
	num   int32
	ref   *oops.Oop_Instance
	_type uint8
}

const (
	_byte    = 1
	_char    = 2
	_double  = 3
	_float   = 4
	_int     = 5
	_long    = 6
	_ref     = 7
	_short   = 8
	_boolean = 9
)
