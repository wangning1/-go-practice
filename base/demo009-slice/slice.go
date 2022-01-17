//切片 动态数组
package main

import (
	"fmt"
)

func main() {
	//slice通过数组

	//slice的创建
	s := make([]int, 3)
	fmt.Println(s)

	s[0] = 1
	s[1] = 2
	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, 4)
	s = append(s, 5)
	s = append(s, 6)
	s = append(s, 7)

	fmt.Println(s)
	fmt.Println(len(s))

	//拷贝到另一个切片中
	c := make([]int, 1)
	copy(c, s)
	fmt.Println(c)
	fmt.Println(len(c))

	//切片操作 slice[low:high]

	l := s[2:]
	fmt.Println("l = ", l)

	m := s[:2]
	fmt.Println("m = ", m)

	n := s[2:4]
	fmt.Println("n = ", n)

	//一行代码中声明并初始化一个 slice 变量
	p := []string{"hello", "world", "golang"}
	fmt.Println(p)

	//多维数据结构
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// 追加一个切片, ... 表示解包，不能省略
	p = append(p, []string{"one", "two"}...)
	fmt.Println(p)

}
