package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
    dummyNode := &ListNode{Next: head}
    
    prefixSums := make(map[int]*ListNode)
    
    currentSum := 0
    current := dummyNode
    for current != nil {
        currentSum += current.Val
        prefixSums[currentSum] = current
        current = current.Next
    }
    
    currentSum = 0
    current = dummyNode
    for current != nil {
        currentSum += current.Val
        current.Next = prefixSums[currentSum].Next
        current = current.Next
    }
    
    return dummyNode.Next
}

func main(){
    head := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: -3,
                Next: &ListNode{
                    Val: 3,
                    Next: &ListNode{
                        Val: 4,
                    },
                },
            },
        },
    }
    node := removeZeroSumSublists(head)
    for node != nil {
        fmt.Println(node.Val)
        node = node.Next
    }
}
