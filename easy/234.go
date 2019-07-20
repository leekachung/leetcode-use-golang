/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil { return true }
    // 声明快慢指针
    fast, slow := head, head
    // 遍历链表 直到slow指针到达链表中点
    for fast != nil && fast.Next != nil {
        // 快指针 每次前进两步
        fast = fast.Next.Next
        // 慢指针 每次前进一步
        slow = slow.Next
    }
    // 反转后半部分链表
    slow = reverseList(slow)
    // 快指针重置
    fast = head
    // 遍历比较
    for slow != nil {
        if fast.Val != slow.Val {
            return false
        }
        fast = fast.Next
        slow = slow.Next
    }
    return true
}

// 反转链表
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
