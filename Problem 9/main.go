package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 100; i < 400; i++ {
		for j := 100; j < 400; j++ {
			var Asq float64 = math.Pow(float64(i), 2)
			var Bsq float64 = math.Pow(float64(j), 2)
			var C float64 = math.Sqrt(Asq + Bsq)
			var sum float64 = float64(i) + float64(j) + C
			var mul float64 = float64(i) * float64(j) * C
			if sum < 1001 && sum > 999 {
				if sum == 1000 {
					fmt.Println(i, "+", j, "=", C, "(+=", sum, ")", "(*=", mul, ")")
				}

			}
		}
	}

}
