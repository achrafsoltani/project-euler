package main

import "fmt"

func main() {
	var n int = 600851475143
	var largestPrimeFactor int = 1
	for k := 2; k < n; k++ {
		if n%k == 0 {
			if isPrime(k) {
				largestPrimeFactor = k
				fmt.Println(largestPrimeFactor)
			}
		}
	}
	// var res bool = isPrime(6)
	fmt.Println(largestPrimeFactor)
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
