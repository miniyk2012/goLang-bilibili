package main

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/thoas/go-funk"

)

func TestPrintSomething(t *testing.T) {
	fmt.Println("Say hi")
	t.Log("Say bye")
	fmt.Println("ab"[2:] == "")
	fmt.Println(len("杨a"))
	if !funk.Contains([]string{"foo", "bar"}, "bar") {
		t.Error("error")
	}
}

func TestSplit(t *testing.T) {
	type test struct {
		name string
		s    string
		sep  string
		want []string
	}
	tests := []test{
		{"simple", "a:b:c", ":", []string{"a", "b", "c"}},
		{"special1", ",", ",", []string{"", ""}},
		{"special2", "a,b,", ",", []string{"a", "b", ""}},
		{"special3", ",b", ",", []string{"", "b"}},
		{"chinese", "杨恺你好", "你", []string{"杨恺", "好"}},
		{"long seq", "abcdefg", "cd", []string{"ab", "efg"}},
		{"long chinese", "杨恺你好呀有往好的呀有看过,你", "呀有", []string{"杨恺你好", "往好的", "看过,你"}},
		{"ahead chinese", "杨恺你好呀有杨", "杨", []string{"", "恺你好呀有", ""}},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := Split(testcase.s, testcase.sep)
			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("excepted:%#v, got:%#v", testcase.want, got)
			}
		})
	}

}
