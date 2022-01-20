// 函数变量
package main

import (
	"fmt"
	"strings"
)

func fire() {
	fmt.Println("fire")
}

//链式处理字符串
func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

//
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	var f func()
	f = fire
	f()

	list := []string{
		"go scnner",
		"go parser",
		"go lang",
		"go printer",
		"go formater",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}

}
