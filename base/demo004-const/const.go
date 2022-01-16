//常量
package main

import (
	"fmt"
	"math"
)

//const 可以出现在任何var出现的地方
const name string = "test"

func main() {
	fmt.Println(name)

	//数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化
	const n = 500000000

	const d = 3e20 / n
	fmt.Println("3e20 / n = ", d)
	fmt.Println("3e20 / n = ", int64(d))

	//一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。
	fmt.Println(math.Sin(n))

}
