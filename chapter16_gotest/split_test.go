package split

import (
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"reflect"
	"testing"
	"time"
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
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
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
		{"all sep", ";;;", ";", []string{"", "", "", ""}},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			tearDownSubTest := setupSubTest(t)
			defer tearDownSubTest(t)
			got := Split(testcase.s, testcase.sep)
			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("excepted:%#v, got:%#v", testcase.want, got)
			}
		})
	}

}

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e:f:g:h:i:j:k:l:::::::::::::n", ":")
	}

}

func TestFib(t *testing.T) {
	tests := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, want := range tests {
		t.Run(fmt.Sprintf("Fib(%d)=%d", i, want), func(t *testing.T) {
			got := Fib(i)
			if got != want {
				t.Errorf("want %d, got %d\n", want, got)
			}
		})
		t.Run(fmt.Sprintf("FastFib(%d)=%d", i, want), func(t *testing.T) {
			got := FastFib(i)
			if got != want {
				t.Errorf("want %d, got %d\n", want, got)
			}
		})
	}

}

// 性能比较函数
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib30(b *testing.B) {
	benchmarkFib(b, 30)
}

func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}

func benchmarkFibCompare(b *testing.B, f func(int) int) {
	for i := 0; i < b.N; i++ {
		f(10)
	}
}

func BenchmarkFibCompareSlow(b *testing.B) {
	benchmarkFibCompare(b, Fib)
}

func BenchmarkFibCompareFast(b *testing.B) {
	benchmarkFibCompare(b, FastFib)
}

func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Printf("write teardown code here...%d\n", retCode) // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(*testing.T) {
	t.Log("测试集前做一些准备工作")
	return func(*testing.T) {
		t.Log("测试集后做一些收尾工作")

	}
}

func setupSubTest(t *testing.T) func(*testing.T) {
	t.Log("子测试前做一些准备工作")
	return func(*testing.T) {
		t.Log("子测试后做一些收尾工作")

	}
}

func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}


