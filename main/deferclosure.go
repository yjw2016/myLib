package main

import "fmt"

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i)
		defer func() { fmt.Println("defer_clousure i=", i) }()
		fs[i] = func() { fmt.Println("clousure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
//匿名函数夺取到i的地址
//defer们在此处开始执行
}
