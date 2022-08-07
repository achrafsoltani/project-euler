package main

import "fmt"
import "strings"
import "strconv"
import "github.com/gammban/numtow"
import "github.com/gammban/numtow/lang"

func main() {
	fmt.Println("hello")

	sum := 0
	for i := 1; i <= 1000; i++ {
		numStr := (numtow.MustString(strconv.Itoa(i), lang.EN))
		numStr = strings.Replace(numStr, "-", "", -1)
		numStr = strings.Replace(numStr, " ", "", -1)
		numStr = strings.Replace(numStr, ",", "", -1)
		fmt.Println(numStr)
		sum += len(numStr)
	}
	fmt.Println(sum)
}
