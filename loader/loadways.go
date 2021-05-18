package loader

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

func LoadingFromZip(fileName, zipPath string) ([]byte, error) {
	// 获取目录下压缩文件的内容
	r, err := zip.OpenReader(zipPath)

	if err != nil {
		return nil, err
	}

	// 关闭该文件描述符
	defer r.Close()

	// 遍历压缩包里面的内容
	for _, f := range r.File {
		// 如果找到了对应类
		if f.Name == fileName {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != err {
				return nil, err
			}
			// 输出所有的数据
			return data, err
		}
	}
	return nil, errors.New("[gvm] not fount: " + zipPath)
}

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
