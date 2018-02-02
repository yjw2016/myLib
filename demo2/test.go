//只是第一个注释
/*
这是一个区块注释
*/
package main

import "fmt"

// 不允许冗余包，不允许冗余变量

func main() {

	var a int = 1
	var b int32 = 2
	a = int(b)

	fmt.Println("使用中文字符")
	fmt.Println(a)

	/*
	  基本类型：
	  int	int8	byte	int16	int32	int64
	  uint8		uint16		uint32    uint64    uint    uintptr

	  复合数据类型：
	  指针: 变量的内存地址  *int *float *string
	  数组: 值类型   传递的是副本    C/C++传的是数组第一个元素指针，golang中传的是副本
	  切片: 可变长度，底层是指针, 2倍扩充
	  结构体: struct
	  接口类型: 非入侵式接口，实现一个接口时，不需要和接口绑定
	*/

}

//心中冷若冰山，如同身处广寒。心中骄阳火辣，恰似油煎火烹。
