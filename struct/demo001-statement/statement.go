//结构体的定义及实例化
package main

import "fmt"

type Point struct {
	X int
	Y int
}

// 同类型变量可以写在一行
type Color struct {
	red, yellow, green int
}

type Commond struct {
	Name    string
	Var     *int
	Comment string
}

func main() {
	// 实例化方式1
	var point Point
	point.X = 20
	point.Y = 30
	fmt.Printf("point %+v\n", point)

	// 实例化方式2
	p := new(Point)
	p.X = 30
	p.Y = 20
	fmt.Printf("p %+v\n", p)

	// 实例化方式3
	var version int = 1
	commond := &Commond{}
	commond.Name = "hello"
	commond.Var = &version
	commond.Comment = "golang"
	fmt.Printf("commond %+v\n", commond)

}
