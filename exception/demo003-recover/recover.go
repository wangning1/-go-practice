// recover 宕机恢复
package main

import (
	"fmt"
	"runtime"
)

type PanicContext struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("运行时异常：", err)
		default:
			fmt.Println("其它异常")
		}
	}()

	entry()
}

func main() {
	fmt.Println("运行前")

	//
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&PanicContext{
			"手动宕机",
		})
		fmt.Println("手动宕机后")
	})

	ProtectRun(func() {
		fmt.Println("复制宕机前")
		var a *int
		*a = 1
		fmt.Println("复制宕机后")
	})
}
