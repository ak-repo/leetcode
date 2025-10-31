/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

    arr1:= []interface{}{}
    arr2:= []interface{}{}
    inorder(p,&arr1)
    inorder(q,&arr2)

    if len(arr1) != len(arr2){
        return false
    }

    for i:= 0; i<len(arr1);i++{

        if arr1[i] != arr2[i]{
            return false
        }
    }

return true
    
   
}



func inorder(node *TreeNode, arr *[]interface{}){

    if node == nil{
    *arr = append(*arr, nil)
        return 
    }
    *arr = append(*arr, node.Val)


    inorder(node.Left,arr)
    inorder(node.Right,arr)

}