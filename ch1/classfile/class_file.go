package classfile

import "fmt"

/*
class文件的映射类
*/
type ClassFile struct {
	// magic uint32
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
	// 字段表,用于表示接口或者类中声明的变量
	fields []*MemberInfo
	// 方法表
	methods []*MemberInfo
	// 属性表
	attributes []AttributeInfo
}

/*
将[]byte解析成ClassFile结构体
*/
func Parse(classData []byte) (cf *ClassFile, err error) {
	// 该部分代码永远都会执行
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
	// start to parse class
	cf.read(cr)
	return
}

/*
read 方法依次调用其他方法解析Class
*/
func (self *ClassFile) read(reader *ClassReader) {
	// 解析魔数
	self.readAndCheckMagic(reader)

	// 解析版本
	self.readAndCheckVersion(reader)

	// 解析常量池
	self.constantPool = readConstantPool(reader)

	// 解析类访问标志
	fmt.Println("[gvm][read] read accessflags ...")
	self.accessFlags = reader.readUint16()

	// 解析本类信息
	fmt.Println("[gvm][read] read class ...")
	self.thisClass = reader.readUint16()

	// 解析父类信息
	fmt.Println("[gvm][read] read superClass ...")
	self.superClass = reader.readUint16()

	// 解析接口
	fmt.Println("[gvm][read] read interfaces ...")
	self.interfaces = reader.readUint16s()

	// 解析字段表
	fmt.Println("[gvm][read] read fields ...")
	self.fields = readMembers(reader, self.constantPool)

	// 解析方法表
	fmt.Println("[gvm][read] read method ...")
	self.methods = readMembers(reader, self.constantPool)

	// 解析属性表
	fmt.Println("[gvm][read] read attribute ...")
	self.attributes = readAttributes(reader, self.constantPool)
}

/*
解析魔术
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	fmt.Println("[gvm][readAdnCheckMagic] read magic ...")
	magic := reader.readUint32()
	// class文件开头是CAFEBABE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
解析版本,主版本号和次版本号都是u2类型
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	fmt.Println("[gvm][readAndCheckVersion] read version ...")
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			fmt.Printf("[gvm][readAndCheckVersion] JDK version is JDK %v.0\n", self.majorVersion-44)
			return
		} else {
			panic("[gvm][readAndCheckVersion] class file version error")
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

/*
解析次版本
*/
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

/*
解析主版本
*/
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

/*
解析常量池
*/
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

/*
解析类访问标志
*/
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/*
在二进制文件中,类名信息存储的是索引,指向了常量池中的位置
*/
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)

}

/*
在二进制文件中,超类信息存储的是索引,指向了常量池中的位置
*/
func (self *ClassFile) SuperClassName() string {
	// if the superClass count more than 1
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

/*
在二进制文件中,接口存储的是索引,指向了常量池中的位置
*/
func (self *ClassFile) InterfaceNames() []string {
	// 切片一个长度位interfaces的string[]数组
	interfaceNames := make([]string, len(self.interfaces))
	// 遍历interfaces,在常量池中查找接口名
	for i, cpIndex := range self.interfaces {
		// 将interfaceName存到interfaceNames中
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	// 返回接口列表
	return interfaceNames
}
