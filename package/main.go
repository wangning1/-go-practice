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
