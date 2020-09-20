package heap

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

func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self Object) Class() *Class {
	return self.class
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
