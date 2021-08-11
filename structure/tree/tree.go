package tree

import (
	"fmt"
)

type Node struct {
	Value string
	Id    int16
}

type BST struct {
	Node

	Left  *BST
	Right *BST
}

type TreeNode struct {
	Value    string
	Id       int16
	ParentId int16
	Children []*TreeNode
}

var MAX_CHILD_ID int = 5
var idCount int16 = 1

func createTree() (treeNode *TreeNode) {
	children := make([]*TreeNode, 0)

	for i := 1; i <= MAX_CHILD_ID; i++ {
		node := &TreeNode{
			Id:       idCount,
			Value:    fmt.Sprint(idCount),
			ParentId: 0,
			Children: make([]*TreeNode, 0),
		}
		idCount++

		children = append(children, node)
	}

	for j, child := range children {
		if j%2 == 0 {
			node := &TreeNode{
				Id:       idCount,
				Value:    fmt.Sprint(idCount),
				ParentId: child.Id,
				Children: make([]*TreeNode, 0),
			}
			idCount++
			child.Children = append(child.Children, node)
		}
	}

	treeNode = &TreeNode{
		Id:       0,
		ParentId: -1,
		Value:    "rt",
		Children: children,
	}

	return treeNode
}

func bfs(node *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		cur := queue[0]
		// shift
		queue = append(queue[:0], queue[1:]...)
		fmt.Printf("%s,", cur.Value)

		for _, child := range cur.Children {
			queue = append(queue, child)
		}
	}
}

func dfs(node *TreeNode) {
	fmt.Printf("%s,", node.Value)

	for _, child := range node.Children {
		dfs(child)
	}
}

func Run() {
	rt := createTree()

	println("tree demo run:")
	print("bfs:")
	bfs(rt)
	println()
	print("dfs:")
	dfs(rt)
	println("\n=======")
}
