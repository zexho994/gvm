package heap

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
	// 参数类型列表
	parameterTypes []string
	// 返回类型
	returnType string
}
