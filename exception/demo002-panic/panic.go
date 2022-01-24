// panic 宕机
package main

import "fmt"

func main() {
	defer fmt.Println("宕机后需处理任务1")
	defer fmt.Println("宕机后需处理任务2")
	panic("手动触发宕机")
}
