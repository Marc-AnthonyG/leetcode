package main

import "fmt"

func main() {
  rootNode := TreeNode{Val: 1}
  test:= TreeNode{Val: 2}
  rootNode2 := TreeNode{1, nil, &test}
	fmt.Println(isSameTree(&rootNode, &rootNode2))
}

// Definition for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
  //Check the current value
	if q == nil || p == nil {
		return q == p
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) 
}

/**
  Note for this problem:

  DFS vs BFS
  DFS will resolve in less loading because we wont have to jump form TreeNode
  of a side to TreeNode of the other side
*/
