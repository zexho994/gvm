package launcher

import "github.com/zouzhihao-994/gvm/src/share/classfile"

// 通过debug模式启动gvm
func StartGvmByDebug() {

	loader := classfile.InitClassLoader(JrePath, UserClassPath)

	if loader == nil {
	}

	loader.Loading("classFile")

}
