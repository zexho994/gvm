package heap

type Object struct {
	// 存放类指针
	class *Class
	// 存放实例变量指针
	fields interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount)}
}
