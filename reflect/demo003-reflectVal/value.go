// 反射的值对象
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 1024
	valueOfA := reflect.ValueOf(a)
	fmt.Printf("valueOfA:%v\n", valueOfA)

	var getA int = valueOfA.Interface().(int)
	var getA2 int = int(valueOfA.Int())
	fmt.Printf("getA:%d, getA2:%d\n", getA, getA2)

}
