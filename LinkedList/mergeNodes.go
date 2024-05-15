package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode{

    curr := head;
    prev := curr;
    next := curr.Next;
    for curr != nil {
        sum := 0
        for curr.Val != 0 && curr != nil{
            sum += curr.Val
            next := curr.Next
            curr = nil;
            curr = next;
        }
        prev.Val = sum;
        prev = next;
        curr = curr.Next
    }
    return head;
}

func main(){
    head := &ListNode{
        Val: 0,
        Next: &ListNode{
            Val: 3,
            Next: &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val: 0,
                    Next: &ListNode{
                        Val: 4,
                        Next: &ListNode{
                            Val: 5,
                            Next: &ListNode{
                                Val: 2,
                                Next: &ListNode{
                                    Val: 0,
                                },
                            },
                        },
                    },
                },
            },
        },
    }
    nodes := mergeNodes(head)
    fmt.Println(nodes)
    //for nodes != nil {
      //  nodes := nodes.Next;
        //fmt.Println(nodes.Val)
    //}
}
