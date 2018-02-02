//接口
/*
interface：
	特征：
    1、接口的实现
    2、接口的组合、嵌套
    3、空接口任意类型
	4、接口断言
	坑：
	当接口包含指针receiver方法时，必须使用指针赋值给接口。
	无法只用值拷贝给接口。
*/
package main

import (
	"fmt"
	"os"
)
type method interface {
	method(string)
}
type ptmethod interface {
	ptmethod(string)
}

//obj嵌套组合了method接口和ptmethod接口(类比java继承)
type obj interface {
	method
	ptmethod
}

//接口的实现，即一个类型实现了一个接口类型所有的方法，即为实现接口，无需显示声明
type person struct {
	name string
}

func (p person) method(str string) {
	fmt.Println(str)
}
func (p *person) ptmethod(str string)  {
	fmt.Println(str)
}

func myPrint(a ...interface{}) (n int, err error)  {
	return fmt.Fprintln(os.Stdout, a...)
}

func main() {

	var p1 person = person{"Jack"}
	var p2 *person = &person{"Tom"}

	//自身，可以调用 自身方法 & 指针方法
	p1.method("自身调用自身方法")
	p1.ptmethod("自身调用指针方法") //此处应该是golang底层给默认了*p1

	//指针，可以调用 自身方法 & 指针方法
	p2.method("指针调用自身方法")
	p2.ptmethod("指针调用指针方法")

	var intf1,intf2 obj

	//坑 intf1 = p1    如果直接把p1赋值给接口类型，报错：
	/*
	报错信息：cannot use p1 (type person) as type obj in assignment:
			person does not implement obj (ptmethod method has pointer receiver)
			也就是说person没有实现obj，而是*person实现了obj
	结论：当接口包含指针receiver方法时，必须使用指针赋值给接口
	*/
	intf1 = &p1
	intf2 = p2

	fmt.Println(intf1,intf2)
	intf1.method("指针接口调用继承而来的method方法")

	//任意类型
	var v1 interface {} = person{"Ada"}
	myPrint(v1)
	v1 = 1
	myPrint(v1)
	v1 = "hello"
	myPrint(v1)
	v1 = map[string]string{"傻子":"小明"}
	myPrint(v1)

	//断言
	if v, bool := v1.(map[string]string) ;bool{
		myPrint(v)
	}else {
		//type 需要和 switch 配套使用
		switch v1.(type){
		case int:
			fmt.Println("v1 为 int")
		default:
			fmt.Println("v1 是什么？who cares")

		}
	}












}
