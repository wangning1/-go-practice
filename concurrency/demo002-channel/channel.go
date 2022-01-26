// 通道
package main

import "fmt"

func Print(c chan int) {
	for {
		data := <-c
		if data == 0 {
			break
		}
		fmt.Println("print data ", data)
	}
	c <- 0
}

func main() {
	// 创建通道
	ch := make(chan interface{})
	go func() {
		fmt.Println("start goroutine")
		ch <- "123"
		fmt.Println("end goroutine")
	}()
	// 通过通道发送和接收数据
	str := <-ch
	fmt.Println("收到通道数据，", str)

	//测试循环发送数据
	c := make(chan int)
	go Print(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}
	c <- 0

	//等待printer结束
	<-c
}
