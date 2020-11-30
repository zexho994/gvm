package launcher

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	jclass2 "github.com/zouzhihao-994/gvm/src/share/jclass"
)

// 通过debug模式启动gvm
func StartGvmByDebug() {

	loader := classfile.InitClassLoader(JrePath, UserClassPath)

	if loader == nil {
	}

	bytecode := loader.Loading("classFile")

	jclass := jclass2.ParseToJClass(bytecode)

	if jclass == nil {

	}

}
