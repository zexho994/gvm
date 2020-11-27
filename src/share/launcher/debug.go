package launcher

import "github.com/zouzhihao-994/gvm/src/share/classfile"

// 通过debug模式启动gvm
func StartGvmByDebug() {

	loader, err := classfile.InitClassLoader(JrePath, UserClassPath)
	if err != nil {
		panic("[gvm] init class load error " + err.Error())
	}

	if loader == nil {

	}

}
