package main

import (
	"fmt"

	tree "github.com/SeoPPak/GoModule"
)

func main() {
	root := tree.InsertNode(50)
	root.Insert(54)
	root.Insert(76)
	root.Insert(45)
	root.Insert(24)
	root.Insert(47)
	root.Insert(94)

	fmt.Println("InOrder (Sorted Order):")
	tree.InOrder(root)
	fmt.Println()
}
