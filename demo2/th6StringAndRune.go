package main

import (
	"fmt"
	"reflect"
)

/*


 */

func main() {
	str := "hello, 91年的岳东卫"
	//遍历的是字符串每个字节
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d " , str[i])
	}

	arr := []byte(str)

	fmt.Println(" ")
	//把字符串转换成字节数组，便利出来的还是每个字节
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d " , arr[i])
	}
	fmt.Println("")
	//遍历每个字符
	for index, code := range str{
		fmt.Printf("index=%d , code=%#q \n", index, code)
	}

	//将字符串转换成rune类型的数组然后遍历
	runeArr := []rune(str)
	fmt.Println("runeArr 的类型是:", reflect.TypeOf(runeArr).Name())
	for index ,code := range runeArr {
		fmt.Printf("index=%d , code=%#q \n", index, code)
	}

	//获取字符串某一个字符
	fmt.Printf("last char is %#q\n", runeArr[len(runeArr)-1])
	fmt.Printf("first char is %#q\n", runeArr[0])
	part := runeArr[0:5]
	fmt.Println(string(part))

















}