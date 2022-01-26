// 通过反射调用方法
package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func main() {
	funcAdd := reflect.ValueOf(add)
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	resultList := funcAdd.Call(paramList)
	fmt.Printf("resultList:%v\n", resultList)
	fmt.Println(resultList[0].Int())
}
