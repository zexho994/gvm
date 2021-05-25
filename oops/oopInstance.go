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

func (i *OopInstance) IsInstanceOf(t *klass.Klass) bool {
	return _checkcast(i.Klass, t)
}

func _checkcast(s, t *klass.Klass) bool {

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !utils.IsInterface(s.AccessFlags) {
			if !utils.IsInterface(t.AccessFlags) {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !utils.IsInterface(s.AccessFlags) {
				return t.IsObject()
			} else {
				return t.IsSuperInterfaceOf(s)
			}
		}
	} else { // s is array
		if !t.IsArray() {
			if !utils.IsInterface(t.AccessFlags) {
				return t.IsObject()
			} else {
				return t.IsJlCloneable() || t.IsJioSerializable()
			}
		} else { // t is array
			//sc := s.GetComponentClass()
			//tc := t.GetComponentClass()
			//return sc == tc || _checkcast(sc, tc)
			panic("_checkcast todo")
		}
	}

	return false
}
