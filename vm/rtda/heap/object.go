package heap

type Object struct {
	// 存放类指针
	class *Class
	// 存放实例变量指针
	// inerface{}可以包含所有类型
	// 对于普通类来说data依然是Slots[]数组
	// 对于数组来说可以是任何类的元素
	data interface{}
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
