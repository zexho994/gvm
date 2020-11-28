package classfile

import (
	"io/ioutil"
	"path/filepath"
)

type ApplicationLoader struct {
	path string
}

func NewApplicationLoader(path string) *ApplicationLoader {
	return &ApplicationLoader{path: path}
}

func (apploader *ApplicationLoader) AddZip(s string) {
	panic("implement me")
}

func (apploader *ApplicationLoader) Path() string {
	panic("implement me")
}

func (apploader *ApplicationLoader) Loading(fileName string) []byte {
	// 拼接目录和类名
	filename := filepath.Join(apploader.path, fileName)

	// 读取目标位置下对应的class文件数据
	data, _ := ioutil.ReadFile(filename)

	// 输出
	return data
}
