/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Val:0,Next:head}
    first,second:=dummy,dummy
    for i:=0;i<n+1;i++ {
        first=first.Next
    }
    for first!=nil {
        first=first.Next
        second=second.Next
    }
    second.Next=second.Next.Next
    return dummy.Next
}