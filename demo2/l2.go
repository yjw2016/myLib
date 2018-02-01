package demo2

import (
	"fmt"
	"runtime"
	domain22 "yhs-go/houtu/domain"
)

func main() {
	f6()
}

func f6() {
	var slice1 = []string{"haha", "hehe", "heihei"}
	fmt.Printf("slice1地址：%p， slice1[0]地址：%p \n", &slice1, &slice1[0])
	slice1[0] = "papapa"
	fmt.Println("修改元素的值之后:")
	fmt.Printf("slice1地址：%p， slice1[0]地址：%p \n", &slice1, &slice1[0])
	slice1 = append(slice1, "hello", "world")
	fmt.Println("经过append之后:")
	fmt.Printf("slice1地址：%p， slice1[0]地址：%p \n", &slice1, &slice1[0])
}

func f5() {
	var a, b, c = "aaa", "bbb", "ccc"
	fmt.Println(a, b, c)
	fmt.Printf("%p,%p \n", &domain22.Class4Age, &domain22.Class4Age[0])
	fmt.Printf("%p,%p \n", &domain22.Class5Age, &domain22.Class5Age[0])
	fmt.Println(domain22.Class4Age)
	fmt.Println(domain22.Class5Age)
	domain22.Class4Age[0] = 1000
	domain22.Class5Age[0] = 3333
	fmt.Println("edit")
	fmt.Printf("%p,%p \n", &domain22.Class4Age, &domain22.Class4Age[0])
	fmt.Printf("%p,%p \n", &domain22.Class5Age, &domain22.Class5Age[0])
	fmt.Println(domain22.Class4Age)
	fmt.Println(domain22.Class5Age)
	domain22.Class4Age[0] = 0
	domain22.Class5Age = append(domain22.Class5Age, 10, 11)
	fmt.Println("edit")
	fmt.Printf("%p,%p \n", &domain22.Class4Age, &domain22.Class4Age[0])
	fmt.Printf("%p,%p \n", &domain22.Class5Age, &domain22.Class5Age[0])
	fmt.Println(domain22.Class4Age)
	fmt.Println(domain22.Class5Age)

}
func f4() {
	stu1 := domain22.Student{"xiaoming", 22, "Man"}
	fmt.Println(stu1, domain22.XiaomingAge)

}

func f3() {
	defer fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}
func f2() {

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
		fallthrough
	case "darwin":
		fmt.Println("OS X")
		fallthrough
	default:
		fmt.Println("你猜")

	}
}
func f1() {
	var a [10]int
	a[0] = 100
	a[1] = 98
	a[3] = 99
	a[4] = 95
	fmt.Println(a[0], a[1], a[2], a[3], a[4])
	fmt.Println(a)
	var b [2]string
	b[0] = "Hello"
	b[1] = "World"
	fmt.Println(b[0], b[1])
	fmt.Println(b)
}
