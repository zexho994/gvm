package classfile

/*
字段表
*/
type MemberInfo struct {
	// 常量池指针
	cp ConstantPool
	// 访问类型
	accessFlags uint16
	// 字段名索引 -> 常量池
	nameIndex uint16
	// 描述符索引 -> 常量池
	descriptorIndex uint16
	// 属性表
	attributesCount uint16
	attributes      []AttributeInfo
}

/*
获取方法的Code属性
*/
func (memberInfo *MemberInfo) CodeAttribute() *CodeAttribute {
	// 遍历属性表
	for _, attrInfo := range memberInfo.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}

	return nil
}

/*
读取字段表
*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 字段的数量
	fieldsCount := reader.readUint16()
	members := make([]*MemberInfo, fieldsCount)

	// 遍历数组
	for i := range members {
		// 解析每一个字段和方法
		members[i] = readMember(reader, cp)
	}

	return members
}

/*
解析字段表数据
*/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	accessFlags := reader.readUint16()
	nameIndex := reader.readUint16()
	descriptorIndex := reader.readUint16()
	attributesCount := reader.readUint16()
	attributes := readAttributes(attributesCount, reader, cp)
	return &MemberInfo{
		cp:              cp,
		accessFlags:     accessFlags,
		nameIndex:       nameIndex,
		descriptorIndex: descriptorIndex,
		attributesCount: attributesCount,
		attributes:      attributes,
	}
}

func (memberInfo *MemberInfo) AccessFlags() uint16 {
	return memberInfo.accessFlags
}

/*
获取方法或字段名称
*/
func (memberInfo *MemberInfo) Name() string {
	return memberInfo.cp.getUtf8(memberInfo.nameIndex)

}

/*
获取字段或方法的描述符
*/
func (memberInfo *MemberInfo) Descriptor() string {
	return memberInfo.cp.getUtf8(memberInfo.descriptorIndex)
}

func (memberInfo *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range memberInfo.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
