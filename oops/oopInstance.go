package oops

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/utils"
)

type OopInstance struct {
	markWords *MarkWords
	*OopFields
	isArray  bool
	isString bool
	jString  string
	*JArray
	*klass.Klass
}

func (o *OopInstance) MarkWord() *MarkWords {
	return o.markWords
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
	isSuper := false
	var f OopField
	for f, isSuper = targetOop.GetField(n, isSuper); true != isSuper; {
		// todo: find from super
	}
	return f, true
}

// NewOopInstance create non-array oops
func NewOopInstance(k *klass.Klass) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		OopFields: InitOopFields(k),
		isArray:   false,
		Klass:     k,
	}
}

// NewArrayOopInstance create array oops
func NewArrayOopInstance(arrayData *JArray) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		isArray:   true,
		JArray:    arrayData,
	}
}

func NewStringOopInstance(str string) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		isString:  true,
		jString:   str,
	}
}
