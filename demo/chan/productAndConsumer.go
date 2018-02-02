/*
注意事项：
	chan,一定是一个协程写，一个协程读，若在一个协程里又写又读，肯定死锁！因为在一个指令发生之后，会停止并等待
    buffer chan 可以在一个协程里又写又读(有范围)
	不能使用	var ch chan int 的形式声明，因为 nil chan无法receive也无法send
	尽量不要在函数里面声明chan，其他goroutine不好接收。
	chan等待没有回复，发生在其他goroutine会阻塞，不影响主程序。发生在主程，deadlock！
	若为主程序设置开关，一定在其他goroutine不会被阻塞的地方写入。
*/
package main

import (
	. "fmt"
	"strconv"
	"time"
)

func main() {
	go produce(ch5)
	go consume(ch5)
	time.Sleep(1*time.Second)
	m2()


	<- sw	//主程序等待produce结束
}

//product
type product struct {
	name      string
	attribute string
}
var ch1 = make(chan int)
var ch2 = make(chan int)
//var ch3 = make(chan int)
var ch5 = make(chan product)
//var ch5 = make(chan product,9) //此处，若ch5为buffer chan,会有何不同
var sw = make(chan bool)

//关于权限，隐式转换，当外部读写chan作为实参传入m1后，ch权限为receive-only，不可send。
//但是等m1返回之后，外部chan依然是receive-send。转换时，权限只能往小，不能往大。
func produce(ch chan<- product){
	println("produce")

	for i := 0; i < 9; i++ {
		ch <- product{"小浣熊",strconv.Itoa(i)}
		println(i)
	}
	sw <- true // 阻塞作用


}

func consume(ch chan product) {
	Println("consume")
	for {
		Println(<-ch)
	}

}

//多路复用
func m2()  {
	select {
	case i := <-ch1:
		Println("ch1:",i)
	case i := <-ch2:
		Println("ch2:",i)
	default:
		Println("ch1 ch2 均无写入")
	}

}