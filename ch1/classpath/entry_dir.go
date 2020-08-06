package classpath

import (
	"io/ioutil"
	filepath "path/filepath"
)

type DirEntry struct {
	/*
		用于存放目录的绝对路径
	*/
	absDir string
}

/*
go中没有构造函数的概念
用newXXX()来代替构造函数
该方法返回一个DirEntry实例
*/
func newDirEntry(path string) *DirEntry {
	// 将路径path转化成绝对路径
	absDir, err := filepath.Abs(path)
	// 如果出现错误,打印
	if err != nil {
		print("[gvm] ")
		println(err)
	}

	return &DirEntry{absDir}
}

/*
读取类的数据
*/
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 拼接目录和类名
	filename := filepath.Join(self.absDir, className)
	// 读取目标位置下对应的class文件数据
	data, err := ioutil.ReadFile(filename)
	// 输出
	return data, self, err
}

/*
返回目录信息
*/
func (self *DirEntry) String() string {
	return self.absDir
}
