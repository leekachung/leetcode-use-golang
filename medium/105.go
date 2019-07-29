/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 { return nil }
    root := preorder[0]
    res := &TreeNode{ Val : root }
    idx := 0
    
    for idx < len(inorder) {
        if inorder[idx] == root {
            break
        }
        idx++
    }
    
    res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
    res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
    
    return res
}
