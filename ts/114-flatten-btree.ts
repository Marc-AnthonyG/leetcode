// Given the root of a binary tree, flatten the tree into a "linked list":
//
// The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
// The "linked list" should be in the same order as a pre-order traversal of the binary tree.



class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

function flatten(root: TreeNode | null): void {
  if (!root) {
    return
  }


  const flattenSubTree = (subTree: TreeNode): void => {
    if (subTree.right) {
      flattenSubTree(subTree.right)
    }

    if (subTree.left) {
      flattenSubTree(subTree.left)

      let left = subTree.left
      const right = subTree.right
      subTree.left = null
      subTree.right = left

      while (left.right) {
          left = left.right
      }

      left.right = right
    }
  }

  flattenSubTree(root)
}

flatten(new TreeNode(1, new TreeNode(2, new TreeNode(3), new TreeNode(4)), new TreeNode(5, null, new TreeNode(6))))
/*
Input
[1,2,5,3,4,null,6]

Output
[1,null,2,null,5,null,6]

Expected
[1,null,2,null,3,null,4,null,5,null,6]
*/
