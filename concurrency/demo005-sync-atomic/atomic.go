//多线程间同步-atomic
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	seq int64
)

func GenID() int64 {
	return atomic.AddInt64(&seq, 1)
}

func main() {
	for i := 0; i < 20; i++ {
		go GenID()
	}
	time.Sleep(time.Second * 10)
	fmt.Println("休眠了10秒钟")
	fmt.Println(GenID())

}
