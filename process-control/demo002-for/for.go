//for的几种使用方法
package main

import "fmt"

func main() {
	// 第一种 接一个表达式
	i := 1
	for i <= 5 {
		fmt.Println("i = ", i)
		i++
	}

	// 第二种 接三个表达j
	for j := 0; j < 5; j++ {
		fmt.Println("j = ", j)
	}

	// 第三种 接一个range表达式<遍历一个可迭代的对象>
	myarr := [...]string{"go", "wold"}
	for index, item := range myarr {
		fmt.Println("index = ", index, ",value = ", item)
	}

	// 第四种 不接表达式
	for {
		fmt.Println("no value")
		break
	}

}
