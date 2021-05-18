package jclass

import (
	"sync"
)

// Perm 等同于方法区的概念
// 专门存储 JClassInstance 对象
type Perm struct {
	Space map[string]*JClassInstance
}

var perm *Perm
var once sync.Once

func GetPerm() *Perm {
	once.Do(func() {
		perm = &Perm{Space: make(map[string]*JClassInstance)}
	})
	return perm
}
