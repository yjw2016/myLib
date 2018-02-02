package main

import "fmt"

var course = map[string][]string{
	"计算机导论": {"计算机组成原理", "计算机网络", "算法", "操作系统"},
	"数学导论":  {"微积分", "线性代数", "概率统计"},
	"医学导论":  {"内科", "外科", "儿科"},
	"方法学导论": {"打开冰箱门原理", "把大象放进去", "关闭冰箱门"},
}

var ii = 1
var slice = []int{6, 7, 8, 9, 0}
var array = []int{1, 2, 3, 4, 5}

func topoSort(m map[string][]string) []string {

	var result []string

	for k, v := range m {

		result = append(result, k)
		for _, lesson := range v {
			result = append(result, lesson)
		}
	}

	return result

}

func main() {
	fmt.Println(topoSort(course))
}
