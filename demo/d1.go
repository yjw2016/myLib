//结构体
package main

import "fmt"

type tip struct {
	x int
}

func (tip *tip) add(i int) *tip {
	tip.x += i
	return tip
}

func (tip tip) sub(i int) tip {
	tip.x -= i
	return tip
}

func main() {
	tip := tip{0}
	tip = *tip.add(100)
	addtip := tip.add(100)
	fmt.Printf("tip.x: %v addtip.x: %v \n", tip.x, addtip.x)
	tip = tip.sub(50)
	tipsub := tip.sub(50)
	fmt.Printf("tip.x: %v subtip.x: %v \n", tip.x, tipsub.x)

}
