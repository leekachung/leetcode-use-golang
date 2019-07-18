/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针法 先让快指针领先n个位置 
// 当快指针到达nil时 慢指针即倒数第n个位置
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := &ListNode{ Next : head }
    fast, slow, step := node, node, 0
    for step < n {
        fast = fast.Next
        step++
    }
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return node.Next
}
