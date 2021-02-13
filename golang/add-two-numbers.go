// https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var firstnode *ListNode
var previousnode *ListNode

// Add two decimal numbers contained in singly-linked lists.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    // Initialize for multiple runs.
    firstnode = nil
    previousnode = nil
    
    // Enumerate both lists at the same time.
    node1 := l1
    node2 := l2
    
    var remainder int
    for node1 != nil || node2 != nil {
        
        // Get digits from current node pointers.
        // If no current node, assume zero digit.
        var digit1, digit2 int
        if( node1 != nil ) {
            digit1 = node1.Val
        }
        if( node2 != nil ) {
            digit2 = node2.Val
        }
        
        // Add the two digits, with previous remainder.
        sum := digit1 + digit2 + remainder

        // Deal with remainder if sum is 10 or more.
        if( sum > 9 ) {
            remainder = 1
            sum = sum % 10
        } else {
            remainder = 0
        }
        
        // Create a node for the summed digit
        // and add it to the output list.
        createDigit(sum)
        
        // Move both pointers to next node.
        if node1 != nil { 
            node1 = node1.Next
        }
        if node2 != nil {
            node2 = node2.Next
        }
    }

    // Tack on a node if there is a leftover remainder.
    if( remainder > 0 ) {
        createDigit(remainder)
    }

    return firstnode
}

// Create a node for a digit and add it to the list.
func createDigit(digit int) {
    currentnode := ListNode {digit, nil}
    if previousnode == nil {
        firstnode = &currentnode
    } else {
        previousnode.Next = &currentnode
    }
    previousnode = &currentnode
}
 
