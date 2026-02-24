package bst_operations

import "fmt"

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(val int) *BSTNode {
	return &BSTNode{Val: val}
}

func Search(root *BSTNode, val int) *BSTNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return Search(root.Left, val)
	}
	return Search(root.Right, val)
}

func Insert(root *BSTNode, val int) *BSTNode {
	if root == nil {
		return NewBSTNode(val)
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
	}
	return root
}

func Remove(root *BSTNode, val int) *BSTNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = Remove(root.Left, val)
	} else if val > root.Val {
		root.Right = Remove(root.Right, val)
	} else {
		// Node found
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// Two children: find in-order successor (smallest in right subtree)
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		root.Val = successor.Val
		root.Right = Remove(root.Right, successor.Val)
	}
	return root
}

func InOrder(root *BSTNode) []int {
	if root == nil {
		return []int{}
	}
	result := InOrder(root.Left)
	result = append(result, root.Val)
	result = append(result, InOrder(root.Right)...)
	return result
}

func Run() {
	var root *BSTNode

	// Insert values
	values := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range values {
		root = Insert(root, v)
	}
	fmt.Println("=== BST After Inserting", values, "===")
	fmt.Println("In-order:", InOrder(root))

	// Search
	fmt.Println("\n=== Search ===")
	if node := Search(root, 4); node != nil {
		fmt.Printf("Found: %d\n", node.Val)
	}
	if node := Search(root, 9); node == nil {
		fmt.Println("Not found: 9")
	}

	// Remove leaf node
	fmt.Println("\n=== Remove 1 (leaf) ===")
	root = Remove(root, 1)
	fmt.Println("In-order:", InOrder(root))

	// Remove node with one child
	fmt.Println("\n=== Remove 6 (one child: none, leaf) ===")
	root = Remove(root, 6)
	fmt.Println("In-order:", InOrder(root))

	// Remove node with two children
	fmt.Println("\n=== Remove 3 (two children) ===")
	root = Remove(root, 3)
	fmt.Println("In-order:", InOrder(root))

	// Remove root
	fmt.Println("\n=== Remove 5 (root) ===")
	root = Remove(root, 5)
	fmt.Println("In-order:", InOrder(root))
}
