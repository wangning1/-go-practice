// 结构体的构造函数
package main

import "fmt"

type Address struct {
	Provice string
	City    string
	ZipCode string
}

func NewAddressByProvice(provice string) *Address {
	return &Address{
		Provice: provice,
	}
}

func main() {
	add := NewAddressByProvice("北京")
	fmt.Printf("add %+v\n", add)
}
