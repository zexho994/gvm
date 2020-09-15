package heap

type TypeDescriptor string

/*
基本类型
*/
func (td TypeDescriptor) IsBaseType() bool { return len(td) == 1 }

/*
void类型
*/
func (td TypeDescriptor) IsVoidType() bool { return td == "V" }

/*
对象类型
*/
func (td TypeDescriptor) IsObjectType() bool { return td[0] == 'L' }

/*
数组类型
*/
func (td TypeDescriptor) IsArrayType() bool { return td[0] == '[' }

/*
long或者double类型
*/
func (td TypeDescriptor) IsLongOrDouble() bool { return td == "J" || td == "D" }

/*
方法的描述符
参考点有两个：
1 方法参数类型与数量
2 返回类型

虽然在虚拟机层面返回类型是可以作为方法的特征之一
但是在java中，返回类型不作为方法重载的特征之一。
原因可能是因为方法即使有返回值，调用者也可以不接收方法的返回值，这会导致编译器无法根据返回类型作为判断依据
*/
type MethodDescriptor struct {
	ParameterTypes []TypeDescriptor
	ReturnType     TypeDescriptor
}

/*
获取参数数量
*/
func (md MethodDescriptor) getParamCount() uint {
	return uint(len(md.ParameterTypes))
}

/*

 */
func (md MethodDescriptor) getParamSlotCount() uint {
	slotCount := md.getParamCount()
	for _, paramType := range md.ParameterTypes {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}
