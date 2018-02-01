package demo2

import "fmt"

func main() {
	dispatch4()
}

func dispatch4() {
	innerFunc()
}

/*
  array： 连续存储、同样类型、线性表、定容、定长、值传递(传递副本，值类型)
*/
func arrarDemo() {
	var arr1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var arr2 = [...]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4}
	arr4 := [...]string{"a", "b", "c"}
	fmt.Println("\n", arr1, "\n", arr2, "\n", arr3, "\n", arr4)
}

/*
  slice： 连续存储、同样类型、线性表、不定容、不定长、引用传递
*/

//func sliceDemo() {
//	var slice1 = []int{1,2,3,4,5,6}
//	slice2 := make([]int, 10, 0)
//	var arr = [...]int{1,2,3,4,5,6,7,8,9}
//	var slice3 = arr[:]
//	var slice4 = arr[0:]
//	var slice5 = arr[1:3]
//	fmt.Println(slice1,slice2,slice3,slice4,slice5)
//}

/*
  数组、切片，当做参数传递
*/
func paramDemo() {
	func1 := func(slice []int) {
		slice[0] = 100
	}
	func2 := func(arr [6]int) {
		arr[0] = 100
	}

	slice1 := []int{1, 2, 3, 4, 5, 6}
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1, "\n", arr)
	func1(slice1)
	func2(arr)
	fmt.Println(slice1, "\n", arr)
	/*  此处arr[0]的值并没有改变，因为arr传递的是副本，而slice传递的才是引用  */
}

//三种遍历
func forRange() {
	//
	var tempArr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//for index, item := range tempArr {
	//	fmt.Println("Index: ", index, " Item : ", item)
	//}
	//
	//for index := range tempArr {
	//	fmt.Println("Index: ", index, " Item : ", tempArr[index])
	//}

	for i := 0; i < len(tempArr); i++ {
		fmt.Println("Index: ", i, " Item : ", tempArr[i])
	}
}

/*
  cap：数组/切片，容量；
  len：数组/切片，长度；
  append：切片，追加；切片在append扩容的时候，容量会翻一翻。
  copy： 切片，复制；
*/
func innerFunc() {
	fmt.Println("-------------  内置函数  ------------")
	var tempArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("arrCap: ", cap(tempArr), "arrLen ", len(tempArr))
	tempArr = append(tempArr, 1, 2, 3)
	fmt.Println("arrCap: ", cap(tempArr), "arrLen ", len(tempArr))
	addSlice := []int{21, 22, 23, 24}
	//给切片添加切片
	tempArr = append(tempArr, addSlice...) //三个点 ... 打散，并且当做参数传递
	fmt.Println("arrCap: ", cap(tempArr), "arrLen ", len(tempArr))
	//给切片添加数组
	addArray := [...]int{1, 2, 3, 4, 5}
	//下面一行 报错 : cannot use addArray (type [5]int) as type []int in append
	//tempArr = append(tempArr,addArray...)
	//正确姿势 ：
	tempArr = append(tempArr, addArray[:]...)
	fmt.Println("arrCap: ", cap(tempArr), "arrLen ", len(tempArr))

	//copy : B切片copy A切片
	bSlice := []int{1, 2, 3, 4, 5, 6}
	aSlice := []int{0, 0, 0}
	copy(aSlice[2:], bSlice)
	fmt.Println(aSlice)
}

func MultiArray() {

	fmt.Println("-------------二维数组--------------")
	//定义一个二维数组
	var twoDimArr1 [5][3]int
	fmt.Println(twoDimArr1)
	//定义一个一维数组 数组每一个成员是[]int空切片
	var twoDimArr2 [5][]int
	twoDimArr2[1] = []int{1, 2, 3}
	fmt.Println(twoDimArr2)
	fmt.Println("-------------三维数组--------------")
	//定义一个三维数组
	var threeDimArr2 [2][3][5]int
	fmt.Println(threeDimArr2)
	//定义一个二维数组,第二维每一项是[]int空切片
	var threeDimArr1 [5][10][]int
	threeDimArr1[0][0] = []int{12, 3, 4, 5}
	fmt.Println(threeDimArr1)
	fmt.Println("-------------数组切片结合--------------")
	//一维切片 数组每一项都是一个[3]int类型数组
	var sliceAndArr [][3]int = [][3]int{[3]int{1, 3, 4}, [3]int{4, 5, 4}}
	fmt.Println(sliceAndArr)
	//二维切片
	var slice [][]int
	fmt.Println(slice)
	slice = [][]int{[]int{1, 2, 3}, []int{6, 6, 6}}
	fmt.Println(slice)
	fmt.Println("-------------三维数组遍历--------------")
	//三维数组遍历 threeDimArr3  定义并且初始化
	var threeDimArr3 [3][3][2]int = [3][3][2]int{[3][2]int{[2]int{1, 2}}}
	for i := 0; i < len(threeDimArr3); i++ {
		for j := 0; j < len(threeDimArr3[i]); j++ {
			for k := 0; k < len(threeDimArr3[i][j]); k++ {
				fmt.Println("threeDimArr3[", i, "][", j, "][", k, "]=", threeDimArr3[i][j][k])
			}
		}
	}
}
