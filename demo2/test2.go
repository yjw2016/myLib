package demo2

import (
	"fmt"
	"math"
	"os"
	"reflect"
)

func main() {
	circleTest()
	fmt.Println("执行完毕")
}

//循环
func circleTest() {

	//利用for实现其他语言中while(true){}的效果
	a := 1
	for {
		if a < 100 {
			fmt.Println(a)
			a++
		} else {
			break
		}
	}
	//传统for循环
	for i := 0; i <= 100; i++ {
		fmt.Println("i=", i)

	}

	j := 0
	for ; j < 100; j++ {
		if j < 100 {
			fmt.Println("j=", j)
		} else {
			break
		}
	}

	//for range 用法，遍历map字典或者数组
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7}
	for index, item := range arr {
		fmt.Println("Index", index, "Item=", item)
	}

	for index := range arr {
		fmt.Println("Index", index, "Item=", arr[index])

	}

	//map 循环
	mp := map[string]string{"zhangsan": "1", "lisi": "2"}
	for k, v := range mp {
		fmt.Println("K=", k, "v", v)
	}

	for k := range mp {
		fmt.Println("K=", k, mp[k])
	}
}

//流程控制
func streamCtr() {
	//一、if 条件不能带小括号，花括号跟随条件
	var a int = 2
	if a > 1 {
		fmt.Println("a>1")
	} else if a < 1 {
		fmt.Println("a<1")
	} else {
		fmt.Println("a==1")
	}

	/*
	  二、switch case
	  case分支不需要break字段，默认跳出，如果继续向下走，显示调用 fallthrough
	  两种用法，1、switch 表达式 case 值 2、switch 空 case 表达条件
	*/

	var b int = 2
	var c int = 2

	switch b {
	case 0:
		{
			fmt.Println("b=0")
		}
	case 1:
		{
			fmt.Println("b=1")
		}
	case 2:
		{
			fmt.Println("b=2")
		}
		fallthrough

	case 7:
		{
			fmt.Println("b=7")
		}
	case 8:
		fmt.Println("哈哈哈哈哈哈哈哈哈哈哈哈哈")
		//goto F8
	default:
		{
			fmt.Println("unknow")
			goto END
		}
	}

	switch {
	case c == 0:
		fmt.Println("c=0")
	case c == 1:
		fmt.Println("c=1")
	case c == 2:
		fmt.Println("c=2")
		fallthrough
	case c == 3 || c == 4 || c == 5 || c == 6:
		fmt.Println("c is 4 or 5 or 6 ")
	case c > 6 && c < 11:
		fmt.Println("6 < c < 11")
	default:
		fmt.Println("c is unknow")

	}
	/*
	  三、goto
	  goto 当前代码执行流程跳转到指定Label
	*/

END:
	fmt.Println("跳转到END标签")

	//F8:{
	//	fmt.Println("jump to b=8")
	//}

}

func arithmetic() {
	//算数运算符： + - * / ++ -- %
	var i int = 0
	fmt.Println(i)
	i = i + 1
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	i = i * 5
	fmt.Println(i)
	i = i % 100
	fmt.Println(i)
	//赋值运算符：= += -= *= /= %= 左移赋值<<= 右移赋值>>= 位逻辑 与&= 异或^= 或|=
	i = 64
	fmt.Println(i)
	i <<= 2
	fmt.Println(i)
	i >>= 3
	fmt.Println(i)
	//8 0x08    0000 1000 ^ 0000 1001
	fmt.Println("8 ^ 9 = ", 8^9)
	//0001 & 0010 = 0000
	fmt.Println("1 & 2 : ", 1&2)
	//0001 | 0010 = 0011
	fmt.Println("1 | 2 : ", 1|2)
	//0001 ^ 0010 = 0011
	fmt.Println("1 ^ 2 : ", 1&2)

}

//复数的使用
func testComplex() {
	//complex64 complex128
	var cmplx1 complex64 = complex(1.1, 2.2)
	var cmplx2 complex64 = complex(3.3, 4.4)

	fmt.Println("复数运算", cmplx1+cmplx2, reflect.TypeOf(cmplx1).Name())
}

//bool类型测试
func testBoolean() {
	//逻辑运算符： && || ！
	//关系预算副： <  >  == >= <=
	var b1 bool = (1 > 2) //flase
	var b2 bool = (1 < 2) //true
	var b3 bool = true
	var b4 bool = true
	var b5 bool = false
	var str = "her"
	var b6 bool = (str == "her")
	fmt.Fprintf(os.Stdout, "a1=%t,a2=%t,a3=%t,a4=%t,a5=%t,a6=%t \n", b1, b2, b3, b4, b5, b6)

	fmt.Println("try2", 2)

}

// 整型测试函数
func testInterger() {

	//有符号
	var a1 int8 = 127
	var a2 int16 = -500

	//无符号

	var b1 byte = 100
	var b2 uint8 = 65
	var b3 uint16 = 65535
	fmt.Fprintf(os.Stdout, "a1=%d,a2=%d,a3=%d \n", a1, a2, 2)
	fmt.Fprintf(os.Stdout, "b1=%d,b2=%d,b3=%d \n", b1, b2, b3)

	//int8赋值给int16，涉及类型转换，a->b   b(a)
	a2 = int16(a1)

	var str string = "hello"
	fmt.Println(str)

	var arr []int = []int{1, 2, 3, 4}
	fmt.Println(arr)

}

func testFloat() {
	//浮点数是一种不精确的表示方法，浮点数1==浮点数2，不准确
	var f2 float32 = 1.2
	var f1 float64 = 1.1

	fmt.Fprintf(os.Stdout, "f1=%.2f,f2=%.4f \n", f1, f2)

	var (
		c2 float32 = 1.11000008
		c1 float32 = 1.11
	)
	fmt.Println("c1==c2:", c1 == c2)
	fmt.Println("c2-c1为", math.Dim(float64(c2), float64(c1))) //居然是0

	if math.Dim(float64(c2), float64(c1)) < 0.0001 {
		fmt.Println("c2等于c1")
	} else {
		fmt.Println("c2大于c1")
	}

}
