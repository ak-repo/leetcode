/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

    arr:= []int{}
    inorder(root,&arr)

    return arr
    
}


func inorder(node *TreeNode, arr *[]int){

    if node == nil{
        return 
    }

    inorder(node.Left,arr)
    *arr = append(*arr, node.Val)
    inorder(node.Right,arr)

}