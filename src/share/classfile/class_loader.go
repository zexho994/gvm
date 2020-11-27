package classfile

// 加载字节码文件
type ClassLoader struct {
	Bytecode []byte
	Bl       *BootStrapLoader
	El       *ExtensionLoader
	Al       *ApplicationLoader
}

func InitClassLoader(jre, cp string) *ClassLoader {
	classLoader := ClassLoader{}

	bl, jrePath := newBootStrapLoader(jre)
	el := NewExtensionLoader(jrePath)
	al := NewApplicationLoader(cp)

	classLoader.Bl = bl
	classLoader.El = el
	classLoader.Al = al

	return &classLoader
}
