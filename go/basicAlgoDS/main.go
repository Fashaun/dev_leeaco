package main

import (
	"fmt"
	"os"

	"basicAlgoDS/array_operations"
	"basicAlgoDS/bst_operations"
	"basicAlgoDS/btree_traversal"
	"basicAlgoDS/bubble_sort"
	"basicAlgoDS/deque_array"
	"basicAlgoDS/dijkstra"
	"basicAlgoDS/dynamic_array_list"
	"basicAlgoDS/go_algods_example"
	"basicAlgoDS/insertion_sort"
	"basicAlgoDS/linked_list"
	"basicAlgoDS/merge_sort"
	"basicAlgoDS/graph_adjacency_list"
	"basicAlgoDS/graph_adjacency_matrix"
	"basicAlgoDS/queue_array"
	"basicAlgoDS/queue_linked_list"
	"basicAlgoDS/selection_sort"
	"basicAlgoDS/stack_array"
	"basicAlgoDS/topological_sort"
	"basicAlgoDS/stack_linked_list"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <example_name>")
		return
	}

	btree_traversal.Run()

	switch os.Args[1] {
	case "array_operations":
		array_operations.Run()
	case "bst_operations":
		bst_operations.Run()
	case "go_algods_example":
		go_algods_example.Run()
	case "btree_traversal":

		/* 初始化二元樹 */
		// 初始化節點
		n1 := btree_traversal.NewTreeNode(1)
		n2 := btree_traversal.NewTreeNode(2)
		n3 := btree_traversal.NewTreeNode(3)
		n4 := btree_traversal.NewTreeNode(4)
		n5 := btree_traversal.NewTreeNode(5)
		// 構建節點之間的引用（指標）
		n1.Left = n2
		n1.Right = n3
		n2.Left = n4
		n2.Right = n5

		btree_traversal.LevelOrder(n1)
	case "bubble_sort":
		bubble_sort.Run()
	case "deque_array":
		deque_array.Run()
	case "dijkstra":
		dijkstra.Run()
	case "dynamic_array_list":
		dynamic_array_list.Run()
	case "insertion_sort":
		insertion_sort.Run()
	case "linked_list":
		linked_list.Run()
	case "merge_sort":
		merge_sort.Run()
	case "graph_adjacency_list":
		graph_adjacency_list.Run()
	case "graph_adjacency_matrix":
		graph_adjacency_matrix.Run()
	case "queue_array":
		queue_array.Run()
	case "queue_linked_list":
		queue_linked_list.Run()
	case "selection_sort":
		selection_sort.Run()
	case "stack_array":
		stack_array.Run()
	case "topological_sort":
		topological_sort.Run()
	case "stack_linked_list":
		stack_linked_list.Run()
	default:
		fmt.Printf("Unknown example: %s\n", os.Args[1])
	}
}
