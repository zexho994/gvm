package classloader

import (
	"io/ioutil"
	"path/filepath"
)

/*
从文件加载加载字节码文件
*/
func loadFromDir(path, className string) ([]byte, error) {
	// 拼接目录和类名
	filename := filepath.Join(path, className)

	// 读取目标位置下对应的class文件数据
	data, err := ioutil.ReadFile(filename)

	// 输出
	return data, err
}
