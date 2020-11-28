package classfile

import "fmt"

// 加载字节码文件
type ClassLoader struct {
	Bytecode []byte
	Bl       *BootStrapLoader
	El       *ExtensionLoader
	Al       *ApplicationLoader
}

func InitClassLoader(jre, cp string) *ClassLoader {
	classLoader := ClassLoader{}

	bl := newBootStrapLoader(jre)
	el := NewExtensionLoader(bl.path)
	al := NewApplicationLoader(cp)

	classLoader.Bl = bl
	classLoader.El = el
	classLoader.Al = al

	return &classLoader
}

func (loader *ClassLoader) Loading(fileName string) {
	// 先判断方法区是否已经存在该class

	fileName = fileName + ".class"
	var data []byte
	// 从启动类加载器中获取bytecode
	if data = loader.Bl.Loading(fileName); data == nil {
		if data = loader.El.Loading(fileName); data == nil {
			if data = loader.Al.Loading(fileName); data == nil {
				panic("class")
			}
		}
	}
	// 解析bytecode

	fmt.Println(data)
}
