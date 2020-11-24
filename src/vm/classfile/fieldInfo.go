package classfile

type FieldInfo struct {
	baseInfo []*MemberInfo
}

func (field FieldInfo) BaseInfo() []*MemberInfo {
	return field.baseInfo
}

func readFieldInfo(reader *ClassReader, cp ConstantPool) FieldInfo {
	// 字段的数量
	fieldCount := reader.readUint16()
	members := make([]*MemberInfo, fieldCount)

	// 遍历数组
	for i := range members {
		// 解析每一个字段和方法
		members[i] = readMember(reader, cp)
	}

	return FieldInfo{baseInfo: members}
}
