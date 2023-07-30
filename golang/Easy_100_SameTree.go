package main

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

  //Check to the left
	if p.Left == nil && q.Left != nil {
		return false
	} else if !isSameTree(p.Left, q.Left) {
		return false
	}


  //Check to the right
	if p.Right == nil {
		return q.Right == nil
	} else if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

/**
  Note for this problem:

  DFS vs BFS
  DFS will resolve in less loading because we wont have to jump form TreeNode
  of a side to TreeNode of the other side
*/
