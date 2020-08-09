package classfile

import "fmt"

type ClassFile struct {
	// 次版本
	minorVersion uint16
	// 主版本
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 类访问标志,表示是类还是接口,public还是private等
	accessFlags uint16
	// 本类
	thisClass uint16
	// 父类
	superClass uint16
	// 接口
	interfaces []uint16
	//
	fields []*MemberInfo
	// 方法
	methods []*MemberInfo
	// 属性
	attributes []AttributeInfo
}

/*
将[]byte解析成ClassFile结构体
*/
func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/*
read 方法依次调用其他读取方法
*/
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/*
解析魔术
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()

	// class文件开头是CAFEBABE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
解析版本
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		fmt.Printf("[gvm][readAndCheckVersion] majorversion is %v\n", self.majorVersion)
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			fmt.Printf("[gvm][readAndCheckVersion] minorVersion is %v\n", self.minorVersion)
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

/*
解析次版本
*/
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
} // getter
/*
解析主版本
*/
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
} // getter

/*
解析常量池
*/
func (self *ClassFile) ConstantPool() ConstantPool {
} // getter

/*
TODO:暂时只获取数据,完整的校验以后做
*/
func (self *ClassFile) AccessFlags() uint16 {

} // getter

func (self *ClassFile) Fields() []*MemberInfo {} // getter

func (self *ClassFile) Methods() []*MemberInfo {} // getter

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)

}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	// 切片一个长度位interfaces的string[]数组
	interfaceNames := make([]string, len(self.interfaces))
	// 遍历interfaces
	for i, cpIndex := range self.interfaces {
		// 将interfaceName存到interfaceNames中
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	// 返回接口列表
	return interfaceNames
}
