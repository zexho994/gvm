package loader

import (
	"os"
)
import "path/filepath"
import "strings"

/*
创建一个通配符类型Entry
将path下的所有jar包文件的路径作为string创建一个zipEntry
最后返回一个path路径下所有jar包构成的zipEntry的数组
后续可以通过readClass直接对数组中的所有zipEntry类型的jar包进行遍历，从而搜索class文件
*/
func newWildcardEntry(path string) CompositeEntry {
	// 获取路径(不包含通配符字符),remove *
	baseDir := path[:len(path)-1]
	var compositeEntry []Entry
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是目录而非文件
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 如果后缀是jar包
		if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") {
			// 创建jar包类型的文件到compositeEntry中
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil

	}
	// 遍历每一个
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
