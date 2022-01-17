//数组
//数组 是一个具有编号且长度固定的元素序列
package main

import "fmt"

func main() {
	//一维数组
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(c)

	//数组长度
	fmt.Println("数组c的长度为：", len(c))

	//多维数组
	d := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(d)

}
