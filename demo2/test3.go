package main

import (
	"fmt"
)

func main() {
	dispatch()
}

func dispatch() {
	//printInfo("雅婷",18, "沈阳")
	//name, age, _ := retInfo("雅婷", 18, "沈阳")
	//fmt.Println("Name: ", name, "Age: ", age)
	//var ids []int = delUses(1,2,3,4,5,6,7)
	//fmt.Println(ids)
	//定义一个短变量来保存匿名函数
	//callback := func(str string) {
	//	fmt.Println(str)
	//}
	//callback := getCallback()
	//匿名函数可以当做函数参数，也可以当做函数返回值来使用
	//testCallBack(callback)

	fibonacci := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci())
	}
}

/*
  1、函数，多返回值
  2、调用并接收函数的多返回值，不需要的返回可以用_下划线符号忽略
  3、固定参数和不定参数，函数中的不定参数会自动转换成数组切片
  4、匿名函数
  5、闭包:  （匿名函数+匿名函数的相关的引用）
*/

//斐波那契
func fibonacci() func() int {
	i, j := 0, 0
	return func() int {
		if i == 0 && j == 0 {
			i = 1
			return 1
		} else if i == 1 && j == 0 {
			j = 1
			return 1
		} else {
			tmp := i
			i += j
			j = tmp
			return i
		}
	}
}

//匿名函数当做函数的返回值来使用
func getCallback() func(string string) {
	return func(str string) {
		fmt.Println(str)
	}
}

//匿名函数，可以当做变量来传值
func testCallBack(callback func(string string)) {
	callback("hello,callback")
}

//不定参数, 接收可以是不定参数, 但是返回不可以是不定返回值
func delUses(uids ...int) (newUserIds []int) {
	fmt.Println("Uids: ", uids)
	return uids
}

//带参数 带返回值
func retInfo(name string, age int, addr string) (string, int, string) {
	return name, age, addr
}

//带参数  不带返回值
func printInfo(name string, age int, addr string) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Address:", addr)
}
