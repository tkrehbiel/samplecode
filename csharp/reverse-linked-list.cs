// https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode ReverseList(ListNode head) {
        // Holds the reversed list's head node.
        ListNode reverseHead = null;
        
        // Enumerate input list while building reversed list.
        ListNode current = head;
        while( current != null ) {
            // Cache next node.
            ListNode nextNode = current.next;
            // New next node is the reversed list's head.
            current.next = reverseHead;
            // Reversed list's head becomes the current node.
            reverseHead = current;
            // Visit next node.
            current = nextNode;
        }

        // Return the reversed list's head.
        return reverseHead;
    }
}