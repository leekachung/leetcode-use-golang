/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    return recursionTree(root, []int{})
}

func recursionTree(root *TreeNode, res []int) []int {
    if root != nil {
        res = recursionTree(root.Left, res)
        res = append(res, root.Val)
        res = recursionTree(root.Right, res)   
    }
    return res
}
