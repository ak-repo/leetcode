/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil{
        return true 
    }

    return isMirror(root.Left,root.Right)
}


func isMirror(left, right *TreeNode)bool{

    if left == nil && right == nil{
        return true 
    }

    if left == nil || right == nil{
        return false 
    }

    if right.Val != left.Val{
        return false 
    }

    return isMirror(left.Left,right.Right) && isMirror(right.Left,left.Right)
}