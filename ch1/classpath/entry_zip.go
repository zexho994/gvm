package classpath

import (
	"archive/zip"
	"path/filepath"
)
import "errors"
import ioutil "io/ioutil"

type ZipEntry struct {
	absPath string
}

/*
构造函数
参考{@see entry_dir.go}的newDirEntry()方法
*/
func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		print("[gvm] ")
		println(err)
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 获取目录下压缩文件的内容
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		print("[gvm] ")
		println(err)
	}
	// 关闭该文件描述符
	defer r.Close()
	// 遍历压缩包里面的内容
	for _, f := range r.File {
		// 如果找到了对应类
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				print("[gvm] ")
				println(err)
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != err {
				print("[gvm] ")
				println(err)
			}
			// 输出所有的数据
			return data, self, err
		}
	}
	return nil, nil, errors.New("[gvm] not fount: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
