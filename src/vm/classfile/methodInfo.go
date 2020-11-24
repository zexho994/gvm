package classfile

type MethodInfo struct {
	baseInfo []*MemberInfo
}

func (methodInfo MethodInfo) BaseInfo() []*MemberInfo {
	return methodInfo.baseInfo
}

func readMethodInfo(methodsCount uint16, reader *ClassReader, cp ConstantPool) MethodInfo {
	// 字段的数量
	members := make([]*MemberInfo, methodsCount)

	// 遍历数组
	for i := range members {
		// 解析每一个字段和方法
		members[i] = readMember(reader, cp)
	}

	return MethodInfo{baseInfo: members}
}
