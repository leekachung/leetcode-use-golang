/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// solution 1 迭代
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

// solution 2 递归
func reverseList(head *ListNode) *ListNode {
    return reverse(nil, head)
}

func reverse(prev, cur *ListNode) *ListNode {
    if cur == nil {
        return prev
    }
    head := reverse(cur, cur.Next)
    cur.Next = prev
    return head
}
