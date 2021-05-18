package oops

import (
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/utils"
)

type OopInstance struct {
	markWords      *MarkWords
	fields         *OopFields
	isArray        bool
	isString       bool
	jString        string
	jArray         *JArray
	jclassInstance *jclass.JClassInstance
}

func (o *OopInstance) MarkWord() *MarkWords {
	return o.markWords
}

func (o *OopInstance) Jclass() *jclass.JClassInstance {
	return o.jclassInstance
}

func (o *OopInstance) ArrayLength() uint32 {
	utils.AssertTrue(o.isArray, "class is not array")
	return o.jArray.length
}

func (o *OopInstance) ArrayData() *JArray {
	return o.jArray
}

func (o *OopInstance) Fields() *OopFields {
	return o.fields
}

func (o *OopInstance) JString() string {
	return o.jString
}

// find the field of oopInstance by field name
// n: field name
func (o *OopInstance) FindField(n string) (OopField, bool) {
	targetOop := o
	isSuper := false
	var f OopField
	for f, isSuper = targetOop.fields.GetField(n, isSuper); true != isSuper; {
		// todo: find from super
	}
	return f, true
}

// create non-array oops
func NewOopInstance(jci *jclass.JClassInstance) *OopInstance {
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
