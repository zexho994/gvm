package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type MonitorEntry struct {
	base.NOP
}

func (m *MonitorEntry) Execute(frame *runtime.Frame) {
	// donothing
}

type MonitorExit struct {
	base.NOP
}

func (m *MonitorExit) Execute(frame *runtime.Frame) {
	// donothing
}
