package oops

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
)

type OopInstance struct {
	markWords      *MarkWords
	fields         *OopFields
	isArray        bool
	isString       bool
	jString        string
	jArray         *JArray
	jclassInstance *jclass.JClass_Instance
}

func (o *OopInstance) MarkWord() *MarkWords {
	return o.markWords
}

func (o *OopInstance) Jclass() *jclass.JClass_Instance {
	return o.jclassInstance
}

func (o *OopInstance) ArrayLength() uint32 {
	exception.AssertTrue(o.isArray, "class is not array")
	return o.jArray.length
}

func (o *OopInstance) ArrayData() *JArray {
	return o.jArray
}

func (o *OopInstance) Fields() *OopFields {
	return o.fields
}

// create non-array oops
func NewOopInstance(jci *jclass.JClass_Instance) *OopInstance {
	return &OopInstance{
		markWords:      NewMarkWords(),
		fields:         InitOopFields(jci),
		isArray:        false,
		jclassInstance: jci,
	}
}

// create array oops
func NewArrayOopInstance(arrayData *JArray) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		isArray:   true,
		jArray:    arrayData,
	}
}

func NewStringOopInstance(str string) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		isString:  true,
		jString:   str,
	}
}
