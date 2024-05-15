package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode{
    current := list1;
    diff := b - a + 1;
    for a != 1 {
        current = current.Next
        a--;
    }
    prev := current;
    for diff != 0 {
        diff--;
        current = current.Next
    }
    prev.Next = list2;
    node2 := list2;
    for node2.Next != nil {
        node2 = node2.Next;
    }
    node2.Next = current.Next; 
    return list1
}

func main(){
    list1 := &ListNode{
        Val: 10,
        Next: &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 13,
                Next: &ListNode{
                    Val: 6,
                    Next: &ListNode{
                        Val: 9,
                        Next: &ListNode{
                            Val: 5,
                        },
                    },
                },
            },
        },
    }
    list2 := &ListNode{
        Val: 10000,
        Next: &ListNode{
            Val: 1000001,
            Next: &ListNode{
                Val: 1000002,
            },
        },
    }
    a, b := 3, 4;
    node := mergeInBetween(list1, a, b, list2)
    for node != nil {
        fmt.Println(node.Val)
        node = node.Next
    }
}
