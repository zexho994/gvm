package oops

/*
java类对象，是所有Clas的父类
*/
type Object struct {
	// 存放类指针
	class *Class
	// 存放实例变量指针
	// inerface{}可以包含所有类型
	// 对于普通类来说data依然是Slots[]数组
	// 对于数组来说可以是任何类的元素
	data interface{}

	//
	extra interface{}
}

func (obj *Object) Extra() interface{} {
	return obj.extra
}
func (obj *Object) SetExtra(extra interface{}) {
	obj.extra = extra
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}

func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}

func (obj Object) Class() *Class {
	return obj.class
}

// reflection
func (obj *Object) GetRefVar(name, descriptor string) *Object {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (obj *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
