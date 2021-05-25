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
	fieldDesc string
	field     []utils.Slot
}

func (f fields) Fields() (string, string, uint8, []utils.Slot) {
	return f.fieldName, f.fieldDesc, f.fieldType, f.field
}

func (sfv *StaticFields) SetStaticField(name string, s []utils.Slot) {
	for idx := range sfv.fields {
		if sfv.fields[idx].fieldName == name {
			sfv.fields[idx].field = s
			return
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
}

func (sfv *StaticFields) AddStaticField(name, desc string, s utils.Slot) {
	newField := fields{field: []utils.Slot{}, fieldName: name, fieldDesc: desc, fieldType: s.Type}
	sfv.fields = append(sfv.fields, newField)
}

func (sfv *StaticFields) GetStaticField(name, desc string) *fields {
	for _, f := range sfv.fields {
		if f.fieldName == name &&
			f.fieldDesc == desc {
			return &f
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
