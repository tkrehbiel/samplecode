// https://leetcode.com/problems/linked-list-cycle/
// Compare/contrast with Go version.

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    
    // Does linked list have a cyclical loop?
    public bool HasCycle(ListNode head) {
        ListNode newListHead = null;
        ListNode newListTail = null;

        // Build a new list while iterating the existing list.
        // If a duplicate node is detected, a cycle exists.
        ListNode node = head;
        while( node != null ) {
            // Cache next pointer because we will change it.
            ListNode next = node.next;
            // Iterate the previous entries in the new list looking for duplicates.
            for( ListNode n = newListHead; n != null; n = n.next ) {
                // If we find a duplicate node, list has cycled,
                // so we can return immediately.
                // Uses Object.Equals() for node equality.
                if( n.Equals( node ) ) return true;
            }
            // Add node to new list.
            // Builds new pointer links to avoid any potential loops.
            if( newListTail != null ) newListTail.next = node;
            newListTail = node;
            if( newListHead == null ) newListHead = node;
            node.next = null;
            node = next;
        }

        // If we get here, the input list
        // did not have any cycles and ended normally.
        return false;
    }
}