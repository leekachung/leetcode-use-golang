/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 || len(inorder) == 0 { return nil }
    root := postorder[len(postorder)-1]
    res := &TreeNode{ Val : root }
    idx := 0
    
    for idx < len(inorder) {
        if inorder[idx] == root {
            break
        }
        idx++
    }
    
    res.Left = buildTree(inorder[:idx], postorder[:idx])
    res.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
    
    return res
}
