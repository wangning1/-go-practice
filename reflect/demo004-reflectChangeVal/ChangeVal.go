// 通过反射修改值
// 值可以修改的条件：可寻址、可导出
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 1024
	// 不能直接修改，不能寻址
	// valueOfA := reflect.ValueOf(a)
	// valueOfA.SetInt(3)
	fmt.Println("a: ", a)

	// 正确的修改港式
	valueOfA := reflect.ValueOf(&a)
	valueOfA = valueOfA.Elem()
	valueOfA.SetInt(3)
	fmt.Println("a修改后: ", a)

	// 通过类型创建类型的实例
	typeOfA := reflect.TypeOf(a)
	aIns := reflect.New(typeOfA)
	fmt.Printf("type:%v, kind:%v\n", aIns.Type(), aIns.Kind())

}
