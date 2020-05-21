/*
  https://leetcode.com/problems/add-two-numbers/
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    current := &ListNode{}
    head := current
    remainder := 0
    for {        
        sum := 0
        if (l1 != nil) {
            sum += (*l1).Val
            l1 = (*l1).Next
        }
        if (l2 != nil) {
            sum += (*l2).Val
            l2 = (*l2).Next
        }
        
        sum += remainder
        (*current).Val = sum % 10
        
        remainder = sum / 10
        if (l1 == nil && l2 == nil) {
            if (remainder > 0) {
                newNode := &ListNode{
                    Val: remainder,
                }
                (*current).Next = newNode
            }
            break
        } else {
            newNode := &ListNode{}
            (*current).Next = newNode
            current = newNode
        }
    }

    return head
}
