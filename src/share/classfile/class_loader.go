package classfile

import (
	"fmt"
	"sync"
)

// 加载字节码文件
type ClassLoader struct {
	Bytecode []byte
	Bl       *BootStrapLoader
	El       *ExtensionLoader
	Al       *ApplicationLoader
}

var BSCLoader *BootStrapLoader
var EXCLoader *ExtensionLoader
var APPLoader *ApplicationLoader
var once sync.Once
var ClaLoader *ClassLoader

func newClassLoader() *ClassLoader {
	return &ClassLoader{}
}

// 初始化类加载器
func InitClassLoader(jre, cp string) *ClassLoader {
	once.Do(func() {
		BSCLoader = newBootStrapLoader(jre)
		EXCLoader = newExtensionLoader(BSCLoader.path)
		APPLoader = newApplicationLoader(cp)
		ClaLoader = newClassLoader()
		ClaLoader.Bl = BSCLoader
		ClaLoader.El = EXCLoader
		ClaLoader.Al = APPLoader
	})

	return ClaLoader
}

// 加载字节码文件到方法区 Perm 中
// 加载顺序依次为 BootStrapLoader 、 ExtensionLoader 、  ApplicationLoader
// 《dynamic class loading in the java virtual machine》 url: https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.18.762&rep=rep1&type=pdf
// @param fileName 类名
func (loader *ClassLoader) Loading(fileName string) []byte {
	// 先判断方法区是否已经存在该class

	fileName = fileName + ".class"

	fmt.Println(fileName)
	var data []byte
	// 从启动类加载器中获取bytecode
	if data = BSCLoader.Loading(fileName); data == nil {
		if data = EXCLoader.Loading(fileName); data == nil {
			if data = APPLoader.Loading(fileName); data == nil {
				panic("class")
			}
		}
	}

	return data
}
