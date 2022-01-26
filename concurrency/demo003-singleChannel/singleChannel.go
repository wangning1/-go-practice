// 单向通道:只能发送的通道类型为chan<-，只能接收的通道类型为<-chan
package main

import "fmt"

func main() {
	ch := make(chan int)
	// 定义只发送通道
	var chSendOnly chan<- int = ch
	var chRecvOnly <-chan int = ch

	go func() {
		chSendOnly <- 123
	}()

	data := <-chRecvOnly
	fmt.Println(data)
}
