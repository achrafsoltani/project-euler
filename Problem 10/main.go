package main

import "fmt"

func main() {
	var sum int = 0
	var max int = 2000000

	for i := 1; i <= max; i++ {
		if isPrime(i) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println(sum)
}

func isPrime(n int) bool {
	var result bool = true
	for k := 2; k < n; k++ {
		if n%k == 0 {
			result = false
		}
	}
	return result
}
