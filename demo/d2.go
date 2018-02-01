//接口
package main

import "fmt"

type obj interface {
	method() string
}

type person struct {
	name string
}

func (p *person) method() string {
	fmt.Println("执行string方法")
	return "ok"
}

func main() {
	var obj obj
	person := person{"Jack"}
	obj = &person
	fmt.Printf("%v \n", obj)

}
