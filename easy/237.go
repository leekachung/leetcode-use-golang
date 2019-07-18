/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    // 复制被删除结点的下一个结点的值
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}
