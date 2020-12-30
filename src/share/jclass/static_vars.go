package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

type StaticFields struct {
	fields []fields
	size   uint32
}

type fields struct {
	fieldName string
	field     *utils.Slot
}

func (f *StaticFields) SetField(name string, s *utils.Slot) {
	f.fields[f.size].fieldName = name
	for idx := range f.fields {
		if f.fields[idx].fieldName == name {
			f.fields[idx].field = s
			return
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
}
