// 方法，结构体绑定方法
// 接收器
package main

import "fmt"

type Property struct {
	value int
}

//接收器类型1 指针类型接收器
func (p *Property) SetValue(val int) {
	p.value = val
}

func (p *Property) GetValue() int {
	return p.value
}

// 接收器类型2 非指针类型接收器
// 当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份。
// 在非指针接收器的方法中可以获取接收器的成员值，但修改后无效

type Point struct {
	X int
	Y int
}

func (p Point) AddPoint(anotherP Point) Point {
	defer func() {
		p.X = 100
	}()
	return Point{p.X + anotherP.X, p.Y + anotherP.Y}
}

func main() {
	var pro Property
	pro.SetValue(12)
	val := pro.GetValue()
	fmt.Println("GetValue = ", val)

	p := &Point{
		20,
		30,
	}

	anotherP := &Point{
		10,
		20,
	}
	fmt.Printf("before add %+v", p)
	point := p.AddPoint(*anotherP)
	fmt.Printf("add %+v", point)
	fmt.Printf("after add %+v", p)

}
