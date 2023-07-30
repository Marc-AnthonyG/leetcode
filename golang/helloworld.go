package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
  rootNode := TreeNode{Val: 1}
  test:= TreeNode{Val: 2}
  rootNode2 := TreeNode{1, nil, &test}
	fmt.Println(isSameTree(&rootNode, &rootNode2))
}
