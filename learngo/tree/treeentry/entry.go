package main

import (
	"fmt"
	"showcase-go/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (t *myTreeNode) postOrder() {
	if t == nil || t.node == nil {
		return
	}
	left := myTreeNode{t.node.Left}
	left.postOrder()
	right := myTreeNode{t.node.Right}
	right.postOrder()
	t.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	fmt.Print("In-order traversal: ")
	root.Traverse()

	fmt.Print("My own post-order traversal: ")

	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
