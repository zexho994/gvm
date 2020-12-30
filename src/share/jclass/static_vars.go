package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

type StaticFieldVars struct {
	fields []fields
	size   uint32
}

type fields struct {
	fieldName string
	field     *utils.Slot
}

func (sfv *StaticFieldVars) SetField(name string, s *utils.Slot) {
	sfv.fields[sfv.size].fieldName = name
	for idx := range sfv.fields {
		if sfv.fields[idx].fieldName == name {
			sfv.fields[idx].field = s
			return
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
}

func NewStaticFieldVars() *StaticFieldVars {
	var fields []fields
	return &StaticFieldVars{
		fields: fields,
		size:   0,
	}
}

func (sfv *StaticFieldVars) AddField(name string, s *utils.Slot) {
	sfv.fields = append(sfv.fields, fields{fieldName: name, field: s})
	sfv.size++
}
