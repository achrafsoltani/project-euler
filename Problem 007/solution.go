package main

import "fmt"

func main() {
	var n int = 2 // 1 is not considered prime
	var count int = 0
	for {
		if isPrime(n) {
			count++
		}
		if count == 10001 {
			fmt.Println(n)
			break
		}
		n++
	}
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
