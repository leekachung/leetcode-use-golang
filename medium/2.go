/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{ 0, nil }
    p := res
    // 是否进位 0为否 1为是
    temp := 0
    
    for l1 != nil || l2 != nil || temp != 0 {
        num1, num2 := 0, 0
        if l1 != nil {
            num1, l1 = l1.Val, l1.Next
        }
        if l2 != nil {
            num2, l2 = l2.Val, l2.Next
        }
        sum := num1 + num2 + temp
        temp = sum / 10
        p.Next = &ListNode{ sum % 10, nil }
        p = p.Next
    }
 
    return res.Next
}
