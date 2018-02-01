package demo2

import (
	"fmt"
	"go/types"
)

func main() {
	a()
	//fmt.Println("return:", a()) // 打印结果为 return: 0
}

func a() types.Object {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return fmt.Println("r")
}
