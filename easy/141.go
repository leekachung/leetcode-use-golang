/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// solution 1
func hasCycle(head *ListNode) bool {
    // 声明快慢指针
    fast, slow := head, head
    // 快指针每次两步 慢指针每次一步
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
	// 快慢指针相遇 即有环
        if ( fast == slow ) { return true }
    }
    return false
}

// solution 2
func hasCycle(head *ListNode) bool {
    hash := make(map[*ListNode]bool)
    for next := head; next != nil; next = next.Next {
        if _, ok := hash[next]; ok {
            return true
        } else {
            hash[next] = true
        }
    }
    return false
}
