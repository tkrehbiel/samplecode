// https://leetcode.com/problems/linked-list-cycle/
// Compare/contrast with C# version.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Does linked list have a cyclical loop?
func hasCycle(head *ListNode) bool {
    var newListHead, newListTail, node *ListNode

    // Build a new list while iterating the existing list.
    // If a duplicate node is detected, a cycle exists.
    node = head
    for node != nil  {
        // Cache next pointer because we will change it.
        next := node.Next
        // Iterate the previous entries in the new list looking for duplicates.
        for n := newListHead; n != nil; n = n.Next {
            // If we find a duplicate node, list has cycled,
            // so we can return immediately.
            if n == node {
                return true
            }
        }
        // Add node to new list.
        // Builds new pointer links to avoid any potential loops.
        if( newListTail != nil ) {
            newListTail.Next = node
        }
        newListTail = node
        if( newListHead == nil ) {
            newListHead = node
        }
        node.Next = nil
        node = next
    }
    return false
}
