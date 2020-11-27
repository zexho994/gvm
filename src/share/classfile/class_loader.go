package classfile

import "path/filepath"

// 加载字节码文件
type ClassLoader struct {
	Bytecode []byte
	Bl       *BootStrapLoader
	El       *ExtensionLoader
	Al       *ApplicationLoader
}

func InitClassLoader(jre, cp string) (*ClassLoader, error) {
	// 初始化启动加载器
	jreDir := getJreDirPath(jre)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	bsl := BootStrapLoader{path: jreLibPath}
	bsl.jars = jars(bsl.path)
	// 初始化扩展加载器

	// 初始化应用加载器

	return &ClassLoader{}, nil
}
