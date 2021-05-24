package klass

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/utils"
)

type StaticFields struct {
	fields []fields
}

type fields struct {
	fieldName string
	fieldType uint8
	field     []utils.Slot
}

func (f fields) Fields() (string, uint8, []utils.Slot) {
	return f.fieldName, f.fieldType, f.field
}

func (sfv *StaticFields) SetField(name string, s []utils.Slot) {
	for idx := range sfv.fields {
		if sfv.fields[idx].fieldName == name {
			sfv.fields[idx].field = s
			return
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
}

func (sfv *StaticFields) AddField(name string, s utils.Slot) {
	newField := fields{field: []utils.Slot{}, fieldName: name, fieldType: s.Type}
	sfv.fields = append(sfv.fields, newField)
}

func (sfv *StaticFields) GetField(name string) *fields {
	for idx := range sfv.fields {
		if sfv.fields[idx].fieldName == name {
			return &sfv.fields[idx]
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
	return nil
}

func NewStaticFieldVars() *StaticFields {
	var fields []fields
	return &StaticFields{
		fields: fields,
	}
}
