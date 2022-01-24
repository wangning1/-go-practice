// 结构体内嵌定义及初始化
// 结构体允许其成员字段在声明时没有字段名而只有类型，这种形式的字段被称为类型内嵌或匿名字段
package main

import "fmt"

type BasicColor struct {
	R, G, B float32
}

type Color struct {
	BasicColor
	Alpha float32
}

// 初始化匿名结构体
type Wheel struct {
	Size int
}

type Car struct {
	Wheel
	Engine struct {
		Power int
		Type  string
	}
}

func main() {
	var c Color
	c.R = 1
	c.G = 1
	c.B = 0
	c.Alpha = 1
	fmt.Printf("c %+v\n", c)

	color := Color{
		BasicColor: BasicColor{
			R: 1,
			G: 1,
			B: 0,
		},
		Alpha: 1,
	}
	fmt.Printf("color %+v\n", color)

	// 初始化匿名结构体
	car := Car{
		Wheel: Wheel{
			18,
		},
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 123,
		},
	}
	fmt.Printf("car %+v\n", car)
}
