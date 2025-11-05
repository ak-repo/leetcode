/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    arr:=[]int{}

    postorderHelper(root,&arr)
    return arr
    
}

    
func postorderHelper(node *TreeNode,arr *[]int){

    if node == nil{
        return 
    }

    postorderHelper(node.Left,arr)
    postorderHelper(node.Right,arr)
    *arr = append(*arr,node.Val)
}