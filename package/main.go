// go module包管理项目通过go run main.go进行执行项目
package main

import (
	"fmt"
	"pack/bar"
	"pack/foo"
	"pack/foo/biz"
)

func main() {
	fmt.Println(foo.Foo())
	fmt.Println(bar.Bar())
	fmt.Println(biz.Biz())

}
