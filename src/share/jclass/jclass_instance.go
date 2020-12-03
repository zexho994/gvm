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
	jci.ThisClass = jci.ConstantPool.GetClassName(jclass.ThisClassIdx)
	// 加载父类
	jci.SuperClass = parseSuper(jclass)
	// 加载接口
	jci.Interfaces = parseInterfaces(jclass)
	// 保存到方法区
	GetPerm().Space[jci.ThisClass] = jci

	return jci
}

// 递归解析父类
func parseSuper(jclass *JClass) *JClass_Instance {
	thisName := jclass.ConstantPool.GetClassName(jclass.ThisClassIdx)
	if thisName == "java/lang/Object" {
		return nil
	}
	// 判断是否存在父类
	superName := jclass.ConstantPool.GetClassName(jclass.SuperClassIdx)
	// 方法区存在该类结构
	perm := GetPerm()
	if supre := perm.Space[superName]; supre != nil {
		return supre
	}
	superBytecode := classfile.ClaLoader.Loading(superName)
	superJClass := ParseToJClass(superBytecode)
	return ParseInstance(superJClass)
}

// 递归解析接口
// TODO：判断父接口是否是接口类,以及其他一些验证流程
func parseInterfaces(jclass *JClass) []*JClass_Instance {
	if jclass.InterfacesCount < 1 {
		return nil
	}

	interfaces := make([]*JClass_Instance, jclass.InterfacesCount)
	for i := range jclass.Interfaces {
		iIdx := jclass.Interfaces[i]
		iName := jclass.ConstantPool.GetClassName(iIdx)
		iInstance := &JClass_Instance{}
		// 如果方法区中已经有直接引用
		if iInstance = perm.Space[iName]; iInstance != nil {
			interfaces[i] = iInstance
			continue
		}
		// 没有的情况，进行接口类的加载
		ibytecode := classfile.ClaLoader.Loading(iName)
		iJClass := ParseToJClass(ibytecode)
		// 接口类型验证
		if !isInterface(iJClass.AccessFlags) {
			panic("[gvm] 接口解析错误 :" + iName + "的父接口对象不为 interface 类型")
		}
		interfaces[i] = ParseInstance(iJClass)
	}

	return interfaces
}
