/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // 创建一个新的链表
    head := &ListNode{}
    // 创建一个指针
    cur := head
    // 遍历比较两条有序链表
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            cur = cur.Next
            l1 = l1.Next
        } else {
            cur.Next = l2
            cur = cur.Next
            l2 = l2.Next
        }
    }
    // 处理边界条件
    if l1 != nil {
        cur.Next = l1
    } else {
        cur.Next = l2
    }
    return head.Next
}
