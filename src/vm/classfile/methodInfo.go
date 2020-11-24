package classfile

type MethodInfo struct {
	baseInfo []*MemberInfo
}

func (methodInfo MethodInfo) BaseInfo() []*MemberInfo {
	return methodInfo.baseInfo
}

func readMethodInfo(reader *ClassReader, cp ConstantPool) MethodInfo {
	// 字段的数量
	methodCount := reader.readUint16()
	members := make([]*MemberInfo, methodCount)

	// 遍历数组
	for i := range members {
		// 解析每一个字段和方法
		members[i] = readMember(reader, cp)
	}

	return MethodInfo{baseInfo: members}
}
