package main

import "fmt"

func Divisors(n int) int {
	var count int
	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(Divisors(5000000))
}
