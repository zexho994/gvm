package classfile

// 加载字节码文件
type ClassLoader struct {
	Bytecode []byte
	Bl       *BootStrapLoader
	El       *ExtensionLoader
	Al       *ApplicationLoader
}

// 初始化类加载器
func InitClassLoader(jre, cp string) *ClassLoader {
	classLoader := ClassLoader{}

	bl := newBootStrapLoader(jre)
	el := newExtensionLoader(bl.path)
	al := newApplicationLoader(cp)

	classLoader.Bl = bl
	classLoader.El = el
	classLoader.Al = al

	return &classLoader
}

// 加载字节码文件到方法区 Perm 中
// 加载顺序依次为 BootStrapLoader 、 ExtensionLoader 、  ApplicationLoader
// 《dynamic class loading in the java virtual machine》 url: https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.18.762&rep=rep1&type=pdf
func (loader *ClassLoader) Loading(fileName string) []byte {
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

	return data
}
