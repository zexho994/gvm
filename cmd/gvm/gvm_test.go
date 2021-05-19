package main

import (
	"github.com/zouzhihao-994/gvm/config"
	"testing"
)

func Test_startJVM(t *testing.T) {
	type args struct {
		className     string
		jrePath       string
		userClassPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// normal test
		{name: "test1", args: args{className: "AlgorithmTest", jrePath: config.JrePathDefault, userClassPath: config.UserClassPathDefault}},
		{name: "test2", args: args{className: "LogicTest", jrePath: config.JrePathDefault, userClassPath: config.UserClassPathDefault}},
		{name: "test3", args: args{className: "PrintFieldsTest", jrePath: config.JrePathDefault, userClassPath: config.UserClassPathDefault}},
		{name: "test4", args: args{className: "PrintStaticTest", jrePath: config.JrePathDefault, userClassPath: config.UserClassPathDefault}},
		// missing parameter
		{name: "test5", args: args{className: "PrintStaticTest", jrePath: "", userClassPath: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
