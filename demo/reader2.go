package main

import (
	"io"
	"fmt"
	"os"
)

type MyReader struct{}

func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<10 && o < 1<<2; i++ {
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}
//实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。
func (mr MyReader) Read(b []byte) (index int, err error)  {
	i := 0
	for ;i < len(b);i++ {
		b[i] = 'A'
	}
	return i, nil
}

func main() {
	Validate(MyReader{})
}
