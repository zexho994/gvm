package classfile

type MemberInfo struct {
	// 常量池指针
	cp ConstantPool
	// 访问类型
	accessFlags uint16
	// 字段名或方法名
	nameIndex uint16
	// 字段或方法描述符
	descriptorIndex uint16
	// 属性表
	attributes []AttributeInfo
}

/*
读取字段表或方法表
*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 读取2字节
	memberCount := reader.readUint16()
	// 切片一个2字节的MemberInfo数组[]
	members := make([]*MemberInfo, memberCount)
	// 遍历数组
	for i := range members {
		// 解析每一个字段和方法
		members[i] = readMember(reader, cp)
	}
	return members
}

/*
读取字段或方法数据
*/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(), // 2字节
		nameIndex:       reader.readUint16(), // 2字节
		descriptorIndex: reader.readUint16(), // 2字节
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {}

/*
获取方法或属性名称
*/
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)

}

/*
获取字段或方法的描述符
*/
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
