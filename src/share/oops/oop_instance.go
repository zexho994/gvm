package oops

import "github.com/zouzhihao-994/gvm/src/share/jclass"

type Oop_Instance struct {
	markWords      *MarkWords
	isArray        bool
	data           *JArray
	extra          *interface{}
	jclassInstance *jclass.JClass_Instance
}

func (o *Oop_Instance) MarkWord() *MarkWords {
	return o.markWords
}

func (o *Oop_Instance) Jclass() *jclass.JClass_Instance {
	return o.jclassInstance
}

func (o *Oop_Instance) ArrayLength() {

}

func (o *Oop_Instance) SetData(d *JArray) {
	o.data = d
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
		data:      arrayData,
	}
}
