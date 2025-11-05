/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    arr:= []int{}
    preOrderHelper(root,&arr)
    return arr


  
    
}

func preOrderHelper(node *TreeNode,arr *[]int) {
	if node == nil {
		return
	}
*arr = append(*arr, node.Val)
	preOrderHelper(node.Left,arr)
	preOrderHelper(node.Right,arr)
}