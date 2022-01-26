// 单元测试
// 要求:命名文件以_test结尾，每个测试实例以Test开头
// 通过go test命令执行，不需要main文件

// 运行全部的测试用例 go test .\hello_test.go
// 运行指定的测试用例 go test -v -run TestB .\hello_test.go
package gotest_test

import "testing"

func TestHello(t *testing.T) {
	t.Log("hello world")
}

func TestA(t *testing.T) {
	t.Log("test A")
}

func TestB(t *testing.T) {
	t.Log("test B")
	t.Fail()
}
