/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

if root == nil{
    return [][]int{}
    }

    coll:= [][]int{}
    queue:= []*TreeNode{root}

    for len(queue)> 0{
       var level []int
       
       lvl_size:= len(queue)

       for i:= 0; i<lvl_size; i++{

        node:= queue[0]
        queue = queue[1:]
        if node == nil {
                continue
            }

        level = append(level, node.Val)

        if node.Left != nil{
            queue = append(queue,node.Left)
        }
        if node.Right != nil{
            queue = append(queue, node.Right)
        }


       }
       coll = append(coll,level)

        
    }

return coll
    
}