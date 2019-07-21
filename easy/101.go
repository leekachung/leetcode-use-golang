/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// solution 1
func isSymmetric(root *TreeNode) bool {
    return recursionTree(root, root)
}

func recursionTree(root *TreeNode, root2 *TreeNode) bool {
    if root == nil && root2 == nil { return true }
    if root == nil || root2 == nil { return false }
    return (root.Val == root2.Val) && recursionTree(root.Left, root2.Right) && recursionTree(root.Right, root2.Left)
}

// solution 2 easy to understand
func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }
    return recursionTree(root.Left, root.Right)
}

func recursionTree(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil { 
        return true 
    } else if left != nil && right != nil && left.Val == right.Val {
        return recursionTree(left.Left, right.Right) && recursionTree(left.Right, right.Left)
    }else {
        return false   
    }
}
