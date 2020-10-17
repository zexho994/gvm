package heap

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/vm/classfile"
	"strings"
)

type Class struct {
	// 类访问标志，public,protect,private
	accessFlags uint16
	// 类完全限定名
	name string
	// 父类完全限定名
	superClassName string
	// 接口完全限定名
	interfaceNames []string
	// 常量池指针
	constantPool *ConstantPool
	// 字段表
	fields []*Field
	// 方法表
	methods []*Method
	// 类加载器
	loader *ClassLoader
	// 父类
	superClass *Class
	// 接口
	interfaces []*Class
	// 实例字段数量
	instanceSlotCount uint
	// 静态字段数量
	staticSlotCount uint
	// 静态变量数组
	staticVars Slots
	// 判断类是否被加载过
	initStarted bool

	// 在Class中可以获取到类对象，为了完成反射
	jClass *Object // java.lang.Class
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

/*
判断class有没有初始化过
*/
func (self *Class) InitStarted() bool {
	return self.initStarted
}

/*
初始化类
*/
func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) Fields() []*Field {
	return self.fields
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self Class) StaticVars() Slots {
	return self.staticVars
}

func (self Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

/*
将字节码信息转化成Class结构
*/
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *Class) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

/**
一个类想要访问另一个类，必须含有
1. 两个类在同一个包下
2. 要访问的类public
*/
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()

}

/**
获取包名
lastIndex : "/" 在 name中出现的最后所以下标
如果name = "java/lang/Object","/"出现的最后位置就是9
*/
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		fmt.Printf("[gvm][getPackgeName] 包名 %v \n", self.name[:i])
		return self.name[:i]
	}
	return ""
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

/*
获取静态方法
*/
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

/*
判断方法的ACC_SUPER是否有被标记
*/
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) Name() string {
	return self.name
}

/*

 */
func (self *Class) GetClinitMethod() *Method { return self.getStaticMethod("<clinit>", "()V") }

/*
返回与类对应的数组类
*/
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
