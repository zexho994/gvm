package launcher

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	jclass "github.com/zouzhihao-994/gvm/src/share/jclass"
)

// 通过debug模式启动gvm
func StartGvmByDebug() {
	loader := classfile.InitClassLoader(JrePath, UserClassPath)
	bytecode := loader.Loading("classFile")
	jClass := jclass.ParseToJClass(bytecode)
	classfile.Linked(jClass)
}
