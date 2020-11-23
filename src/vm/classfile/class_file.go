package classfile

import "fmt"

/*
class文件的映射类
*/
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
	// 字段表,用于表示接口或者类中声明的变量
	fields []*MemberInfo
	// 方法表
	methods []*MemberInfo
	// 属性表
	attributes []AttributeInfo
}

// 将字节码的二进制数据[]byte解析成ClassFile结构体
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

	// start to parse class
	cf.read(cr)
	return
}

/*
read 方法依次调用其他方法解析Class
*/
func (classFile *ClassFile) read(reader *ClassReader) {
	// 解析魔数
	classFile.readAndCheckMagic(reader)

	// 解析版本
	classFile.readAndCheckVersion(reader)

	// 解析常量池
	classFile.constantPool = readConstantPool(reader)

	// 解析类访问标志
	//fmt.Println("[gvm][read] read accessflags ...")
	classFile.accessFlags = reader.readUint16()

	// 解析本类信息
	//fmt.Println("[gvm][read] read class ...")
	classFile.thisClass = reader.readUint16()

	// 解析父类信息
	//fmt.Println("[gvm][read] read superClass ...")
	classFile.superClass = reader.readUint16()

	// 解析接口
	//fmt.Println("[gvm][read] read interfaces ...")
	classFile.interfaces = reader.readUint16s()

	// 解析字段表
	//fmt.Println("[gvm][read] read fields ...")
	classFile.fields = readMembers(reader, classFile.constantPool)

	// 解析方法表
	//fmt.Println("[gvm][read] read method ...")
	classFile.methods = readMembers(reader, classFile.constantPool)

	// 解析属性表
	//fmt.Println("[gvm][read] read attribute ...")
	classFile.attributes = readAttributes(reader, classFile.constantPool)

}

/*
解析魔术
*/
func (classFile *ClassFile) readAndCheckMagic(reader *ClassReader) {
	//fmt.Println("[gvm][readAdnCheckMagic] read magic ...")
	magic := reader.readUint32()
	// class文件开头是CAFEBABE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
解析版本,主版本号和次版本号都是u2类型
*/
func (classFile *ClassFile) readAndCheckVersion(reader *ClassReader) {
	//fmt.Println("[gvm][readAndCheckVersion] read version ...")
	classFile.minorVersion = reader.readUint16()
	classFile.majorVersion = reader.readUint16()
	switch classFile.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if classFile.minorVersion == 0 {
			//fmt.Printf("[gvm][readAndCheckVersion] JDK version is JDK %v.0\n", classFile.majorVersion-44)
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
func (classFile *ClassFile) MinorVersion() uint16 {
	return classFile.minorVersion
}

/*
解析主版本
*/
func (classFile *ClassFile) MajorVersion() uint16 {
	return classFile.majorVersion
}

/*
解析常量池
*/
func (classFile *ClassFile) ConstantPool() ConstantPool {
	return classFile.constantPool
}

/*
解析类访问标志
*/
func (classFile *ClassFile) AccessFlags() uint16 {
	return classFile.accessFlags
}

func (classFile *ClassFile) Fields() []*MemberInfo {
	return classFile.fields
}

func (classFile *ClassFile) Methods() []*MemberInfo {
	return classFile.methods
}

/*
在二进制文件中,类名信息存储的是索引,指向了常量池中的位置
*/
func (classFile *ClassFile) ClassName() string {
	return classFile.constantPool.getClassName(classFile.thisClass)

}

/*
在二进制文件中,超类信息存储的是索引,指向了常量池中的位置
*/
func (classFile *ClassFile) SuperClassName() string {
	// if the superClass count more than 1
	if classFile.superClass > 0 {
		return classFile.constantPool.getClassName(classFile.superClass)
	}
	return ""
}

/*
在二进制文件中,接口存储的是索引,指向了常量池中的位置
*/
func (classFile *ClassFile) InterfaceNames() []string {
	// 切片一个长度位interfaces的string[]数组
	interfaceNames := make([]string, len(classFile.interfaces))
	// 遍历interfaces,在常量池中查找接口名
	for i, cpIndex := range classFile.interfaces {
		// 将interfaceName存到interfaceNames中
		interfaceNames[i] = classFile.constantPool.getClassName(cpIndex)
	}
	// 返回接口列表
	return interfaceNames
}
