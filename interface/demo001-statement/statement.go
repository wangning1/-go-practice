// 接口的声明
package main

import (
	"fmt"

	sqlite3 "github.com/mattn/go-sqlite3"
)

//实现接口的两个条件
// 1.接口的类型与实现接口的类型方法格式一致
// 2.接口中所有方法均被实现

type DataWriter interface {
	WriteData(data interface{}) error

	CanWrite() bool
}

type file struct {
}

//接口实现
func (f *file) WriteData(data interface{}) error {
	// 模拟写入
	fmt.Println("write data ", data)
	return nil
}

func main() {
	f := file{}
	f.WriteData("hello")
	libVersion, libVersionNumber, sourceId := sqlite3.Version()
	fmt.Printf("libVersion:%v, libVersionNumber:%v,sourceId:%v", libVersion, libVersionNumber, sourceId)
}
