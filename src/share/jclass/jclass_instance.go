package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/attribute"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

// 存储在方法区中的对象，也是 JClass 经过链接步骤后得到的对象
// 同一个的类或接口的所有子类/实现类对该部分的依赖都会是同一个对象，即不会存在两个一样的 JClass_Instance 对象
type JClass_Instance struct {
	// 常量池
	ConstantPool constant_pool.ConstantPool
	// 类访问标志,表示是类还是接口,public还是private等
	AccessFlags uint16
	// 本类
	ThisClass string
	// 父类
	SuperClass *JClass_Instance
	// 接口
	Interfaces []*JClass_Instance
	// 字段表,用于表示接口或者类中声明的变量
	Fields Fields
	// 方法表
	Methods Methods
	// 属性表
	Attributes []attribute.AttributeInfo
}

func ParseInstance(jclass *JClass) *JClass_Instance {
	jci := &JClass_Instance{}
	// 运行时常量池不变
	jci.ConstantPool = jclass.ConstantPool
	// 类访问符不变
	jci.AccessFlags = jclass.AccessFlags
	// 获取到全限定名
	jci.ThisClass = jci.ConstantPool.GetClassName(jclass.ThisClass)
	// 加载父类
	jci.SuperClass = parseSuper(jclass)
	// 加载接口
	jci.Interfaces = parseInterfaces()

	// 保存到方法区
	GetPerm().Space[jci.ThisClass] = jci

	return jci
}

// 递归解析父类
func parseSuper(jclass *JClass) *JClass_Instance {
	// 判断是否存在父类
	superName := jclass.ConstantPool.GetClassName(jclass.SuperClass)
	//if superName == "java/lang/Object"{
	//	return nil
	//}

	perm := GetPerm()
	// 方法区存在该类结构
	if supre := perm.Space[superName]; supre != nil {
		return supre
	}
	superBytecode := classfile.ClaLoader.Loading(superName)
	superJClass := ParseToJClass(superBytecode)
	return ParseInstance(superJClass)
}

// 递归解析接口
func parseInterfaces() []*JClass_Instance {

	return nil
}
