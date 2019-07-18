/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// solution 1 less code 快慢指针法 先让快指针领先n个位置 
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

// solution 2 遍历两次 不推荐
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if ( head.Next == nil ) {
        return nil
    }
    node := &ListNode{ Next : head }
    pointer, pointer2, length:= node, node, 1
    for pointer.Next != nil {
        pointer = pointer.Next
        length++
    }
    index := length - n
    for i := 1; i <= length; i++ {
        if i == index {
            if pointer2.Next == nil { break }
            pointer2.Next = pointer2.Next.Next
            break
        } else {
            pointer2 = pointer2.Next
        }
    }
    return node.Next
}
