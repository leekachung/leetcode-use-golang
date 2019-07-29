/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// solution 1 4ms
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
    // 是否需要反转 偶数倍反转
	needReverse := true
	nodes := []*TreeNode{root}
    
	for len(nodes) != 0 {
		result = append(result, extractVal(nodes))
		nodes = levelNodes(nodes, needReverse)
		needReverse = !needReverse
	}
	return result
}

func extractVal(in []*TreeNode) []int {
	val := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		val[i] = in[i].Val
	}
	return val
}

func levelNodes(in []*TreeNode, needReverse bool) (result []*TreeNode) {
	if needReverse {
		for i := len(in)-1; i >=0;i--{
			node := in[i]
			if node.Right != nil {
				result = append(result, node.Right)
			}
			if node.Left != nil {
				result = append(result, node.Left)
			}
		}
    } else {
        for i := len(in)-1; i >=0;i--{
            node := in[i]
            if node.Left != nil {
                result = append(result, node.Left)
            }
            if node.Right != nil {
                result = append(result, node.Right)
            }
        }
    }
	return
}

// solution 2 0ms
func zigzagLevelOrder(root *TreeNode) [][]int {
    var res [][]int
    loopprint(root,0,&res)
    return res
}

func loopprint(node *TreeNode,level int,res *[][]int) {
    if node == nil {
        return
    }
    if len(*res) < level+1 {
        *res = append(*res,[]int{node.Val})
    }else {
        if level % 2 == 0 {
            (*res)[level] = append((*res)[level],node.Val)
        }else {
            var tmp []int
            tmp = append(tmp,node.Val)
            tmp = append(tmp, (*res)[level]...)
            (*res)[level] = tmp 
        }
       
    }
    loopprint(node.Left,level+1,res)
    loopprint(node.Right,level+1,res)   
}
