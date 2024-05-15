package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    curr := head

    for curr != nil && curr.Next != nil {
        if curr.Val == curr.Next.Val {
            if curr.Next.Next == nil {
                curr.Next = nil;
            } else {
                curr.Next = curr.Next.Next;
            }
        } else {
            curr = curr.Next
        }
    }
    return head;
}

func main(){
    head := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 2,
                Next: &ListNode{
                    Val: 3,
                    Next: &ListNode{
                        Val: 3,
                    },
                },
            },
        },
    }
    curr := deleteDuplicates(head)
    for curr != nil {
        fmt.Println(curr.Val)
        curr = curr.Next
    }
}
