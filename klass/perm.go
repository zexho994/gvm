package klass

import (
	"sync"
)

// perm 等同于方法区的概念
// 专门存储 Klass 对象
type perm struct {
	Space map[string]*Klass
}

var p *perm
var once sync.Once

func Perm() *perm {
	once.Do(func() {
		p = &perm{Space: make(map[string]*Klass)}
	})
	return p
}
