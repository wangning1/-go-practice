// 空接口
package main

import "fmt"

// 类型断言 x.(T) 其实就是判断 T 是否实现了 x 接口，
// 如果实现了，就把 x 接口类型具体化为 T 类型；
// 而 x.(type) 这种方式的类型断言，就只能和 switch 搭配使用，因为它需要和多种类型比较判断，以确定其具体类型。
func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int 类型")
	case string:
		fmt.Println("string 类型")
	case bool:
		fmt.Println("bool 类型")
	}
}

func main() {
	var a interface{} = 1
	var b interface{} = "hello"
	fmt.Println("a == b : ", a == b)

	//不可以比较空接口的动态值,比较将触发宕机
	// var c interface{} = []int{10}
	// var d interface{} = []int{20}
	// fmt.Println("c == d : ", c == d)

	printType(1)
	printType("hello")
	printType(true)

}
