/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    ch := make(chan int, k)
    var walk func(node *TreeNode)
    walk = func(node *TreeNode){
        if node == nil{
            return
        }
        walk (node.Left)
        
        ch <- node.Val

        walk (node.Right)
    }

    go walk(root)
    
    res := 0
    for i := 0; i < k; i++ {
        res = <-ch
    }
    
    return res
}
