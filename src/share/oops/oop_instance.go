package oops

import "github.com/zouzhihao-994/gvm/src/share/jclass"

type Oop_Instance struct {
	markWords      *MarkWords
	data           interface{}
	extra          interface{}
	jclassInstance *jclass.JClass_Instance
}

func NewOopInstance(jci *jclass.JClass_Instance) *Oop_Instance {
	return &Oop_Instance{
		markWords:      NewMarkWords(),
		jclassInstance: jci,
	}
}
