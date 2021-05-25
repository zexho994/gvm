package oops

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/utils"
)

type OopInstance struct {
	*MarkWords
	*OopFields
	isArray  bool
	isString bool
	jString  string
	*JArray
	*klass.Klass
}

func (o *OopInstance) ArrayLength() uint32 {
	utils.AssertTrue(o.isArray, "class is not array")
	return o.JArray.length
}

func (o *OopInstance) JString() string {
	return o.jString
}

// FindField find the field of oopInstance by field name
// n: field name
func (o *OopInstance) FindField(n string) (OopField, bool) {
	targetOop := o
	var f OopField
	var isSuper bool
	for f, isSuper = targetOop.GetField(n); isSuper == false; {
		// todo: find from super
	}
	return f, true
}

// NewOopInstance create non-array oops
func NewOopInstance(k *klass.Klass) *OopInstance {
	return &OopInstance{
		MarkWords: NewMarkWords(),
		OopFields: InitOopFields(k),
		isArray:   false,
		Klass:     k,
	}
}

// NewArrayOopInstance create array oops
func NewArrayOopInstance(arrayData *JArray) *OopInstance {
	return &OopInstance{
		MarkWords: NewMarkWords(),
		isArray:   true,
		JArray:    arrayData,
	}
}

func NewStringOopInstance(str string) *OopInstance {
	return &OopInstance{
		MarkWords: NewMarkWords(),
		isString:  true,
		jString:   str,
	}
}
