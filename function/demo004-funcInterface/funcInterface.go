//函数实现接口
package main

import "fmt"

//声明接口
type Invoker interface {
	Call(interface{})
}

//声明结构体
type Struct struct {
}

//结构体实现函数
func (s *Struct) Call(p interface{}) {
	//调用实现
	fmt.Println("struct call param ", p)
}

//函数实现接口
type FuncCaller func(interface{})

//实现 Invoker
func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.Call("hello")

	//函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体。
	//当类型方法被调用时，还需要调用函数本体。
	var interfaceInvoker Invoker
	interfaceInvoker = FuncCaller(func(i interface{}) {
		fmt.Println("interfaceInvoker call ", i)
	})
	interfaceInvoker.Call("golang")

}
