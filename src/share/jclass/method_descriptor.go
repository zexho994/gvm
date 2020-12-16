package jclass

import (
	"strings"
)

type MethodDescriptor struct {
	parameteTypes string
	returnType    string
}

// descriptor => (...parameteTypes)...returnType
// for instance: void fun(int) -> (I)V or ; public int method3(Object obj) -> (Ljava/lang/Object)I
func ParseMethodDescriptor(method MethodInfo) MethodDescriptor {
	idx := method.descriptorIdx
	descStr := method.CP().GetUtf8(idx)
	splits := strings.Split(descStr, ")")
	parametes := strings.Split(splits[0], "(")[0]
	return MethodDescriptor{
		parameteTypes: parametes,
		returnType:    splits[1],
	}
}
