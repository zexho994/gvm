package main

import (
	"github.com/zouzhihao-994/gvm/congifuration"
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
		{name: "test1", args: args{className: "AlgorithmTest", jrePath: congifuration.JrePath, userClassPath: congifuration.UserClassPath}},
		{name: "test2", args: args{className: "LogicTest", jrePath: congifuration.JrePath, userClassPath: congifuration.UserClassPath}},
		{name: "test3", args: args{className: "PrintFieldsTest", jrePath: congifuration.JrePath, userClassPath: congifuration.UserClassPath}},
		{name: "test4", args: args{className: "PrintStaticTest", jrePath: congifuration.JrePath, userClassPath: congifuration.UserClassPath}},

		// missing parameter
		{name: "test5", args: args{className: "PrintStaticTest", jrePath: "", userClassPath: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startJVM(tt.args.className, tt.args.jrePath, tt.args.userClassPath)
		})
	}
}
