// 通过反射获取结构体的成员类型
package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
	// 带有结构体tag字段
	Type int `json:"type" id:"100"`
}

func main() {
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfIns := reflect.TypeOf(ins)

	// 遍历结构体所有成员
	for i := 0; i < typeOfIns.NumField(); i++ {
		filedType := typeOfIns.Field(i)
		fmt.Printf("Name:%v,tag:%v\n", filedType.Name, filedType.Tag)
	}

	// 通过字段名称找反射类型
	if catType, ok := typeOfIns.FieldByName("Type"); ok {
		fmt.Printf("tag json = %s, id = %s\n", catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
