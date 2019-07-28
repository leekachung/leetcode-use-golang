/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }
    // 奇偶指针
    odd, even := head, head.Next
    for even != nil && even.Next != nil {
        temp := odd.Next
        odd.Next = even.Next
        even.Next = even.Next.Next
        odd.Next.Next = temp
        
        odd = odd.Next
        even = even.Next
    }
    return head
}
