package classpath

import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	// 获取路径(不包含通配符字符),remove *
	baseDir := path[:len(path)-1]

	compositeEntry := []Entry{}
	//
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
			jarEntry := newEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil

	}
	// 遍历每一个
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
