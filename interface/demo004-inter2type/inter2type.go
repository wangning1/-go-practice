// 接口和类型互转
package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird fly")
}

func (b *bird) Walk() {
	fmt.Println("bird walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, obj := range animals {
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)

		// 判断董伟是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name is %s is Flyer : %v ,is Walker : %v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}

	// 将接口转换为类型
	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)
	fmt.Printf("p1 = %p, p2 = %p", p1, p2)

	// 将a转换为*bird类型将运行时错误
	// p3 := a.(*bird)
	// fmt.Printf("p3 = %p", p3)
}
