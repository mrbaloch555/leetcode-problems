package main

import "fmt"

func winnerOfGame(colors string) bool {
	// res := false
	return finder(colors, "A", "Alice", 0, 0)
}

func finder(colors string, target string, turn string, Acount int, Bcount int) bool {
	for i := 1; i < len(colors); i++ {
		if len(colors[i:]) > 2 {
			if turn == "Alice" {
				if string(colors[i-1]) == target && string(colors[i]) == target && string(colors[i+1]) == target {
					colors = colors[0:i] + colors[i+1:]
					Acount++
					finder(colors, "B", "Bob", Acount, Bcount)
				}
			} else {
				if string(colors[i-1]) == target && string(colors[i]) == target && string(colors[i+1]) == target {
					colors = colors[0:i] + colors[i+1:]
					Bcount++
					finder(colors, "A", "Alice", Acount, Bcount)
				}
			}
		}
	}
	fmt.Println(Acount, Bcount)
	return false
}

func main() {
	colors := "AAAAABBBB"
	fmt.Println(winnerOfGame(colors))
}
