// 闭包 = 函数 + 引用环境
package main

//闭包的作用
// 1.闭包可以访问函数内部的局部变量
// 2.闭包由于每个实例引用的变量都是同一个地址，所以闭包函数里面的局部变量，
// 可以当成全局变量来进行使用，达到全局变量的效果，并且每个实例对应的局部变量就是一个新的全局变量，
// 把需要“记忆功能”的变量都封装到了函数中去。
import "fmt"

func f1() func() int {
	counter := 0
	f2 := func() int {
		counter++
		return counter
	}
	return f2
}

func main() {
	a := f1()
	fmt.Println(a())
	fmt.Println(a())

	b := f1()
	fmt.Println(b())
}
