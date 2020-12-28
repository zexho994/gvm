package oops

import "github.com/zouzhihao-994/gvm/src/share/jclass"

type Oop_Instance struct {
	markWords      *MarkWords
	data           interface{}
	extra          interface{}
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

func NewOopInstance(jci *jclass.JClass_Instance) *Oop_Instance {
	return &Oop_Instance{
		markWords:      NewMarkWords(),
		jclassInstance: jci,
	}
}
