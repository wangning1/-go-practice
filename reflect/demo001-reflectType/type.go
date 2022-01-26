// 反射类型对象
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	type cat struct {
	}

	//获取结构体实例的反射对象
	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	// 指针和指针指向的元素
	ins := &cat{}
	typeOfC := reflect.TypeOf(ins)
	fmt.Printf("name:'%v';kind:'%v'\n", typeOfC.Name(), typeOfC.Kind())

	// Go程序中对指针获取反射对象时，可以通过reflect.Elem()方法获取这个指针指向的元素类型。
	// 这个获取过程被称为取元素，等效于对指针类型变量做了一个“*”操作
	typeOfC = typeOfC.Elem()
	fmt.Printf("name:'%v';kind:'%v'\n", typeOfC.Name(), typeOfC.Kind())

}
