/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    // Previous node
    var prev *ListNode = nil

    // Current node
    curr := head

    // Traverse the linked list
    for curr != nil {

        // Store next node
        nextNode := curr.Next

        // Reverse pointer
        curr.Next = prev

        // Move prev forward
        prev = curr

        // Move curr forward
        curr = nextNode
    }

    // prev becomes new head
    return prev
}