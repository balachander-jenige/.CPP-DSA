/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

    curr := head

    for curr != nil && curr.Next != nil {

        if curr.Val == curr.Next.Val {
            // remove duplicate
            curr.Next = curr.Next.Next
        } else {
            // only move when no duplicate
            curr = curr.Next
        }
    }

    return head
}