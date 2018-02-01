package main

import (
	"fmt"
)

func main() {
	sli := []int{1, 3, 7, 2, 4, 8, 6, 0, 5, 9}
	var tree *tree
	for _, v := range sli {
		tree = AddTree(tree, v)
	}
	fmt.Println(tree)
	sli2 := []int{}
	sli2 = SortTree(tree, sli2)
	fmt.Println(sli2)

}

//中序插入
func AddTree(t *tree, v int) *tree {

	if t == nil {
		t = new(tree)
		t.data = v

	} else if t.data < v {
		t.rTree = AddTree(t.rTree, v)

	} else {
		t.lTree = AddTree(t.lTree, v)

	}

	return t

}

//中序遍历
func SortTree(tree *tree, result []int) []int {

	if tree.lTree != nil {
		result = SortTree(tree.lTree, result)

	}

	result = append(result, tree.data)

	if tree.rTree != nil {
		result = SortTree(tree.rTree, result)
	}
	return result

}

type tree struct {
	data         int
	lTree, rTree *tree
}
