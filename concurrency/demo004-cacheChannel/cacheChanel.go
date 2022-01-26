// 带缓冲的通道
package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	//带缓冲区的通道，当通道被填满时将会发生阻塞
	// ch <- 4
	fmt.Println("ch len", len(ch))
	close(ch)
	for data := range ch {
		fmt.Println("data:", data)
	}

}
