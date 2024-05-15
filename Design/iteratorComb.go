package main

import (
	"fmt"
)

type CombinationIterator struct {
	characters   string
	combinations []string
	currentIndex int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	var combinations []string
	generateCombinations(&combinations, "", characters, combinationLength, 0)
	return CombinationIterator{
		characters:   characters,
		combinations: combinations,
		currentIndex: 0,
	}
}

func generateCombinations(combinations *[]string, currentCombination, characters string, combinationLength, start int) {
	if len(currentCombination) == combinationLength {
		*combinations = append(*combinations, currentCombination)
		return
	}

	for i := start; i < len(characters); i++ {
		generateCombinations(combinations, currentCombination+string(characters[i]), characters, combinationLength, i+1)
	}
}

func (this *CombinationIterator) Next() string {
	if this.currentIndex < len(this.combinations) {
		nextCombination := this.combinations[this.currentIndex]
		this.currentIndex++
		return nextCombination
	}
	return ""
}

func (this *CombinationIterator) HasNext() bool {
	return this.currentIndex < len(this.combinations)
}

func main() {
	iterator := Constructor("abc", 2)
	fmt.Println(iterator.Next())    // Output: "ab"
	fmt.Println(iterator.HasNext()) // Output: true
	fmt.Println(iterator.Next())    // Output: "ac"
	fmt.Println(iterator.HasNext()) // Output: true
	fmt.Println(iterator.Next())    // Output: "bc"
	fmt.Println(iterator.HasNext()) // Output: false
}
