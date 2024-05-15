package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode){
    slow, fast := head, head;

    for fast != nil && fast.Next != nil {
        slow = slow.Next;
        fast = fast.Next.Next;
    }

    reversed := reverse(slow.Next);
    slow.Next = nil;
    current := head
	for current != nil && reversed != nil {
		nextCurr := current.Next
		nextRev := reversed.Next

		current.Next = reversed
		reversed.Next = nextCurr

		current = nextCurr
		reversed = nextRev
	}    
}
func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    reversed := reverse(head.Next)
    newNode := &ListNode{Val: head.Val}
    current := reversed
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode

    return reversed
}
func main(){
    head := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val: 5,
                    },
                },
            },
        },
    }
    reorderList(head)
    for head != nil {
        fmt.Println(head.Val);
        head = head.Next
    }
}
