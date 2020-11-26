package oops

/*
判断是否是数组类型
*/
func (class *Class) IsArray() bool {
	return class.name[0] == '['
}

/*
根据count创建数组对象
*/
func (class *Class) NewArray(count uint) *Object {
	if !class.IsArray() {
		panic("Not array class: " + class.name)
	}
	switch class.Name() {
	case "[Z":
		return &Object{class, make([]int8, count), nil}
	case "[B":
		return &Object{class, make([]int8, count), nil}
	case "[C":
		return &Object{class, make([]uint16, count), nil}
	case "[S":
		return &Object{class, make([]int16, count), nil}
	case "[I":
		return &Object{class, make([]int32, count), nil}
	case "[J":
		return &Object{class, make([]int64, count), nil}
	case "[F":
		return &Object{class, make([]float32, count), nil}
	case "[D":
		return &Object{class, make([]float64, count), nil}
	default:
		return &Object{class, make([]*Object, count), nil}
	}
}

func (class *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(class.name)
	return class.loader.LoadClass(componentClassName)
}
