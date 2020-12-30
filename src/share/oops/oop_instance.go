package oops

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
)

type Oop_Instance struct {
	markWords      *MarkWords
	isArray        bool
	jArray         *JArray
	jclassInstance *jclass.JClass_Instance
}

func (o *Oop_Instance) MarkWord() *MarkWords {
	return o.markWords
}

func (o *Oop_Instance) Jclass() *jclass.JClass_Instance {
	return o.jclassInstance
}

func (o *Oop_Instance) ArrayLength() uint32 {
	exception.AssertTrue(o.isArray, "class is not array")
	return o.jArray.length
}

func (o *Oop_Instance) ArrayData() *JArray {
	return o.jArray
}

// create non-array oops
func NewOopInstance(jci *jclass.JClass_Instance) *Oop_Instance {
	return &Oop_Instance{
		markWords:      NewMarkWords(),
		isArray:        false,
		jclassInstance: jci,
	}
}

// create array oops
func NewArrayOopInstance(arrayData *JArray) *Oop_Instance {
	return &Oop_Instance{
		markWords: NewMarkWords(),
		isArray:   true,
		jArray:    arrayData,
	}
}
