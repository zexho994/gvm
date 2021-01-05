package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/utils"
)

type StaticFieldVars struct {
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

func (sfv *StaticFieldVars) SetField(name string, s []utils.Slot) {
	for idx := range sfv.fields {
		if sfv.fields[idx].fieldName == name {
			sfv.fields[idx].field = s
			return
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
}

func (sfv *StaticFieldVars) AddField(name string, s utils.Slot) {
	newField := fields{field: []utils.Slot{}, fieldName: name, fieldType: s.Type}
	sfv.fields = append(sfv.fields, newField)
}

func (sfv *StaticFieldVars) GetField(name string) *fields {
	for idx := range sfv.fields {
		if sfv.fields[idx].fieldName == name {
			return &sfv.fields[idx]
		}
	}
	exception.GvmError{Msg: "not found static field"}.Throw()
	return nil
}

func NewStaticFieldVars() *StaticFieldVars {
	var fields []fields
	return &StaticFieldVars{
		fields: fields,
	}
}
