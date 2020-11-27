package classfile

import (
	"os"
	"path/filepath"
	"strings"
)

type Loader interface {
	Loading()
	AddZip(string)
	Path() string
}

/*
创建一个通配符类型Entry
将path下的所有jar包文件的路径作为string创建一个zipEntry
最后返回一个path路径下所有jar包构成的zipEntry的数组
后续可以通过readClass直接对数组中的所有zipEntry类型的jar包进行遍历，从而搜索class文件
*/
func jars(path string) []string {
	// 获取路径(不包含通配符字符),remove *
	baseDir := path[:len(path)-1]
	var jars []string
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
			jars = append(jars, path)
		}
		return nil

	}
	// 遍历每一个
	filepath.Walk(baseDir, walkFn)
	return jars
}

/*
查找jre目录的路径
*/
func getJreDirPath(jreOption string) string {
	// 如果用户输入了-Xjre 参数
	if jreOption != "" && exists(jreOption) {
		//fmt.Printf("[gvm][getJreDir] 获取到/jre的路径 : %v\n", jreOption)
		return jreOption
	}
	/*
		如果用户没有输入 -Xjre 参数
		在 './jre' 下找
	*/
	if exists("./jre") {
		//fmt.Printf("[gvm][getJreDir] 获取到/jre的路径\n")
		return "./jre"
	}

	/**
	没有输入 -Xjre参数,而且当前目录下也没有找到
	在 JAVA_HOME中找
	*/
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		//fmt.Printf("[gvm][getJreDir] find JAVA_HOME dir")
		return filepath.Join(jh, "jre")
	}

	// 3种情况都不存在jre目录,输出错误
	panic("[gvm][getJreDir]Can't find jre folder")
}

/*
判断path目录是否存在
*/
func exists(path string) bool {
	// stat获取项目中path的文件信息
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
