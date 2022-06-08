package main

import "fmt"

func main() {
	var rank int = 1; 
	for {
		var triangle int = generate_triangle_number(rank)
		var divisors int = count_divisors(triangle)
		fmt.Println(triangle, " - ", divisors)
		if divisors > 500 {
			fmt.Println("Found =>", triangle)
			break
		}
		rank++
	}
}

func count_divisors(n int) int {
	var divisors int = 0
	for i := 1; i <= n; i++ {
		if (n % i == 0){
			divisors++
		}
	}
	return divisors
}

func generate_triangle_number(rank int) int {
	var triangle int = 0
	for i := 1; i <= rank; i++{
		triangle += i
	}
	return triangle
}