package main

import (
	"fmt"
	"math"
)
func findDivisors(n int) []int {
    divisors := []int{1}
    if n == 0 {
        return divisors
    }

    sqrtN := int(math.Sqrt(float64(n)))

    for i := 2; i <= sqrtN; i++ {
        if n%i == 0 {
            if n/i == i {
                // If i is a perfect square divisor, add it only once
                divisors = append(divisors, i)
            } else {
                // Add both i and n/i
                divisors = append(divisors, i, n/i)
            }
        }
    }

    return divisors
}

func sum(nums []int) int {
    sum := 0;
    for _, num := range nums {
        sum += num;
    }
    return sum;
}
func isOdd(sum int, n int) bool {
    fmt.Println(n, sum, sum % 2 == 1)
    return sum %2 == 1;
}

func isPrime(n int) bool {
    if n <= 1 {
        return false;
    }
    if n <= 3 {
        return true;
    }

    if n%2 == 0 || n%3 == 0 {
        return false;
    }

    for i := 5; i*i <= n; i += 6 {
        if n%i == 0 || n%(i+2) == 0 {
            return false;
        }
    }
    return true;
}

func perfectNumber(){
    n := 1;
    sum := 0;
    for true {
        sum += n;
        if isPrime(sum){
            // fmt.Println("IS Prime: ", sum)
            fmt.Println("Perfect number: ", (sum * n))
        }
        n += n;
    }
}
func main(){
    // n := 1;
    // for n < 1000000 {
    //     divisors := findDivisors(n)
    //     sum := sum(divisors)
    //     if sum == n {
    //         isOdd(n, sum)
    //     }
    //     n++;
    // }
    perfectNumber();
}
