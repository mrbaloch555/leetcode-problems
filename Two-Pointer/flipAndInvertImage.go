package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
    for i := 0; i < len(image); i++ {
        l, r := 0, len(image[i]) - 1;
        for l < r {
            image[i][r], image[i][l] = image[i][l], image[i][r]
            l++;
            r--;
        }
        for j := 0; j < len(image[i]); j++{
            image[i][j] = image[i][j] ^ 1;
        }
    }
    return image
}
func main(){
    image := [][]int{{1,1,0}, {1,0,1}, {0,0,0}}
    fmt.Println(flipAndInvertImage(image))
}
