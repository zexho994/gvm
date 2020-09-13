package heap

type Object struct {
	// 存放类指针
	class *Class
	// 存放实例变量指针
	//
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) Fields() Slots {
	return self.fields
}
