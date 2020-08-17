package classpath

import (
	"os"
	"strings"
)

/*
系统的路径列表分隔符
*/
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {

	/*
		负责寻找和加载Class文件
		文件路径是相对路径,文件名有后缀,比如 java/lang/Object.class
	*/
	readClass(className string) ([]byte, Entry, error)

	/*
		相当Java中的toString()方法
	*/
	String() string
}

/**
根据参数创建不同的Entry实例
*/
func newEntry(path string) Entry {
	// 列表类型 "/"
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 通配符类型 "*"
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// 压缩包类型
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	// 返回目录形式的类路径
	return newDirEntry(path)

}
