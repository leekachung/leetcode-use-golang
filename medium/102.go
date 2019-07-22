/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// solution 1 bfs
func levelOrder(root *TreeNode) [][]int {
    return bfs(root)
}
func bfs(root *TreeNode) [][]int {
    if root == nil { return [][]int{} }
    // 声明队列
    queue := make([] *TreeNode, 0)
    // 添加元素
    queue = append(queue, root)
    // 声明返回值
    levels := make([][]int, 0)
    
    // len(queue) 按照题目所给参数为1而不是7
    for len(queue) > 0 {
        n, level := len(queue), make([]int, 0)、
        // 遍历每层的元素 并追加不为空的左右子树
        for i := 0; i < n; i++ {
            // 取出队列头
            root = queue[0]
            // 清空队列
            queue = queue[1:]
            // 追加每层的值
            level = append(level, root.Val)
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right != nil {
                queue = append(queue, root.Right)
            }
        }
        // 合并每层的值
        levels = append(levels, level)
    }
    return levels
}

// solution 2
func levelOrder(root *TreeNode) [][]int {
    return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) <= level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
    return res
}
