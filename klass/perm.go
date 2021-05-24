package klass

import (
	"sync"
)

type space map[string]*Klass

// perm 等同于方法区的概念
// 专门存储 Klass 对象
type perm struct {
	space
}

var Perm *perm
var once sync.Once

func InitPerm() {
	once.Do(func() {
		Perm = &perm{make(map[string]*Klass)}
	})
}

func PermSpace() map[string]*Klass {
	return Perm.space
}

func (p *perm) Save(name string, klass *Klass) {
	p.space[name] = klass
}

func (p *perm) Get(name string) *Klass {
	return p.space[name]
}
