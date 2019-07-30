/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// solution 1 20ms use channel
func kthSmallest(root *TreeNode, k int) int {
    c := make(chan int)
    go func() {
        Traverse(c, root)
        defer close(c)
    }()
    x, ok := <- c
    for k > 1 && ok {
        x, ok = <- c
        k--
    }
    return x
}

func Traverse(c chan int, root *TreeNode) {
    if root == nil { return }
    Traverse(c, root.Left)
    c <- root.Val
    Traverse(c, root.Right)
}

// solution 2 12ms
func kthSmallest(root *TreeNode, k int) int {
	kv := 0
	var findKth func(r *TreeNode)
	findKth = func(r *TreeNode) {
		if r == nil || k == 0 {
			return
		}
		findKth(r.Left)
		k--
		if k == 0 {
			kv = r.Val
			return
		}
		findKth(r.Right)
	}
	findKth(root)
	return kv
}
