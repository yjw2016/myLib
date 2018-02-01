package demo2

import (
	"fmt"
	"log"
	//"math/rand"
	"os"
	//"time"
)

var c, python, java bool
var i, j = 1, 2

func main() {

	//fmt.Println("main函数执行，时间：", time.Now())
	//fmt.Println("My random number is", rand.Intn(190))
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println("My random number is", r.Intn(190))
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println("My random number is", r.Intn(190))
	//
	//fmt.Println(isay(), add(2, 5))
	//fmt.Println(swap("world", "Hello"))
	//fmt.Println(split(197))
	//
	//x := 3
	//y := "I am Joel Smith."
	//var i = 10000
	//fmt.Println(i, c, python, java, x, y)

	i, j := 45, 1973
	p := &i
	//p保存的是i的地址
	fmt.Println(*p)
	//*p是指针，通过这个指针去读取i
	*p = 27
	//赋值27给指针*p
	fmt.Println(i)
	//打印i
	p = &j
	//p保存的地址换成了j的地址
	*p = *p / 37
	//用指针带入j做除法
	fmt.Println(j)
	//打印j

}

//init函数会在main之前执行
func init() {
	log.SetOutput(os.Stdout)
	fmt.Println("init被执行")
}

//有两个参数的函数
func add(x, y int) int {
	return x + y
}

//没有参数的函数
func isay() string {
	return "I say the number is"
}

//多个返回值的函数
func swap(x, y string) (string, string, string) {
	return y, x, "!!!"
}

//命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
