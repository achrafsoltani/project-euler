package main

import "fmt"
import "math"
import "strconv"

func main() {
	x := math.Pow(2, 1000)
	digits := fmt.Sprintf("%f", x)
	var sum int = 0
	for _, val := range digits {
		num, _ := strconv.Atoi(string(val))
		sum += int(num)
	}
	fmt.Println("sum : ", sum)
}
