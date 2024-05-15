package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}
    fmt.Println(fast)
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func findDuplicateII(nums []int) []int {
    duplicates := make([]int, 0)

    for _, num := range nums {
        index :=  abs(num) - 1;
        fmt.Println(num, index)
        if nums[index] < 0 {
            duplicates = append(duplicates, abs(num))
        } else {
            nums[index] = -nums[index]
        }
    }
    return duplicates
}

func abs(x int) int {
    if x < 0 {
        return -x;
    }
    return x;
}

func main() {
	nums := []int{2, 1, 3, 1, 4, 5, 2}
	fmt.Println(findDuplicate(nums))
    fmt.Println(findDuplicateII(nums))
}
