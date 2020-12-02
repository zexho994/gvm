package launcher

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
)

// 通过debug模式启动gvm
func StartGvmByDebug() {
	loader := classfile.InitClassLoader(JrePath, UserClassPath)
	bytecode := loader.Loading("classFile")
	jclass.ParseToJClass(bytecode)
}
