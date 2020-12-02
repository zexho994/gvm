package jclass

import (
	"sync"
)

// 等同于方法区的概念
// 专门存储 JClass_Instance 对象
type Perm struct {
	Space map[string]*JClass_Instance
}

var perm *Perm
var once sync.Once

func GetPerm() *Perm {
	once.Do(func() {
		perm = &Perm{Space: make(map[string]*JClass_Instance)}
	})
	return perm
}
