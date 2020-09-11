package heap

import (
	"../../classfile"
	"../../classpath"
	"fmt"
)

/*
类加载器
类加载器依赖ClassPath来搜寻和读取class文件
*/
type ClassLoader struct {
	cp *classpath.Classpath

	// 已经加载的类，key是类的全限定名
	classMap map[string]*Class //loaded classed
}

/*
创建一个加载器实例
*/
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

/*
在classMap中根据name查询类
然后将将类数据加载到方法区中
*/
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}

/*
非数组类的加载
*/
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	// 类的链接
	link(class)
	fmt.Printf("[LOADED %s from %s]\n", name, entry)
	return class
}

/*
在classpath中搜索名称为name的类
*/
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}

/*
根据类数据获得类结构体
*/
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 将类的数据转换成类结构体
	class := parseClass(data)
	// 设置类的加载器
	class.loader = self
	// 解析父类以及接口
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

/*
将类数据解析成类结构体
*/
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

/*
解析超类的符号引用
如果父类不是Object类，就获取父类的名称然后调用类加载器加载父类
*/
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

/*
解析接口的符号引用
获取接口名数组 -> 对每一个父接口进行加载
*/
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	fmt.Printf("[gvm][verify] 类加载-验证阶段")
}

// 准备阶段
func prepare(class *Class) {
	fmt.Printf("[gvm][verify] 类加载-准备阶段")
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

/**
计算实例字段数量
1. 父类的字段都属于字段。子类的字段表需要加上父类的字段
2.
*/
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

/**
给类变量分配空间，然后赋予初始值
*/
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalvar(class, field)
		}
	}
}

/**
给类变量赋值
类变量的值在编译时候就已知，所以可以直接从class文件常量池中获取
*/
func initStaticFinalvar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.slotId
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
