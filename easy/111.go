// 不能使用math.min() 空节点不计入 即a, b 当a=0 || b =0时 返回另外一个数 而不是0
func Min(a int, b int) int {
    if a==0 {               // 小坑，空节点不计入最小深度的计算
        return b
    } else if b==0 {
        return a
    }
    if a<b {                // 正常的比较
        return a
    }
    return b
}
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return Min(minDepth(root.Left), minDepth(root.Right)) + 1   // 返回当前节点左右子树的最小深度
}
