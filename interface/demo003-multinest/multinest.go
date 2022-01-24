// 接口的嵌套组合
package main

import "io"

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type device struct {
}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func main() {
	var wc io.WriteCloser = new(device)
	wc.Write(nil)
	wc.Close()

	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)
}
