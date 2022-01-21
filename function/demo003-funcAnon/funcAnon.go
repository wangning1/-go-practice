// 匿名函数
package main

import (
	"fmt"
	"strings"
)

//匿名函数作为回调函数
func visit(list []string, f func(string)) {
	for _, str := range list {
		f(str)
	}
}

func main() {
	//匿名函数声明
	func(data int) {
		fmt.Println("anonymous func print", data)
	}(100)

	//匿名函数复制给变量
	f := func(name string) {
		fmt.Println("hello ", name)
	}
	f("王宁")

	//匿名函数作回调函数使用
	list := []string{"happy", "new", "year"}
	visit(list, func(s string) {
		strNew := strings.ToUpper(s)
		fmt.Println(strNew)
	})

}
