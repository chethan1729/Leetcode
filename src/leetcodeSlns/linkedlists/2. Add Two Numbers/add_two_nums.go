/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	
	// head will keep the reference to the head of the result node so that it can returned back later
	head := &ListNode{}
	
	// prev is used to parse the linked list and to perfom sum operations b/w linked lists
	// sum/=10 will take carry from previous sum and adds it to next
    for prev,sum := head,0; l1!=nil || l2!=nil || sum !=0; sum/=10 {
        
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
		
		// sum%10 will neglect carry and assins the remainder as the value
        prev.Next = &ListNode{
            Val : sum%10,
        }
        prev = prev.Next
    }  
    
    return head.Next
}