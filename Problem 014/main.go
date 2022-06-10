package main

import "fmt"

func main() {
	var selected, largest int = 0, 0
	for s := 1; s < 1000000; s++ {
		count := count_sequence(s)
		if count > largest {
			largest = count 
			selected = s
		}
		fmt.Println("S: ", s)
	}
	fmt.Println("Largest", largest)
	fmt.Println("Selected", selected)
}

func count_sequence(n int) int {
	var count int = 1
	for n != 1 {
		if n%2 == 0 {
			n = n/2
		} else {
			n = 3*n +1
		}
		count++
	}
	return count
}