package classpath

import (
	"errors"
	"strings"
)

/*
多个Entry组合
*/
type CompositeEntry []Entry

/*
构造函数
将参数根据'/'转化成若干Entry,然后组成CompositeEntry
*/
func newCompositeEntry(pathList string) CompositeEntry {
	//fmt.Printf("[gvm][newCompositeEntry] create compositeEntry <pathList> : %v \n", pathList)
	// 创建CompositeEntry对象
	var compositeEntry []Entry
	// 将pathList根据"/"切片后,遍历
	for _, path := range strings.Split(pathList, pathListSeparator) {
		// 根据每一个path创建entry
		entry := newEntry(path)
		// entry加入到list中
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry

}

/**
读取类的方法
依次调用子entry的readClass方法,读取到就返回
*/
func (self CompositeEntry) readClass(ClassName string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(ClassName)

		// 如果找到了
		if err == nil {
			//fmt.Printf("[gvm][compositeEntry.readClass] read %v success\n", ClassName)
			return data, from, nil
		}
	}
	return nil, nil, errors.New("[gvm][readClass] class not found : " + ClassName)
}

/*
toString() 方法
调用每一个子类entry的String()即可
*/
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	// i自增
	for i, entry := range self {
		strs[i] = entry.String()
	}
	// strs拼接起来,中间放置pathListSeparator
	return strings.Join(strs, pathListSeparator)

}
