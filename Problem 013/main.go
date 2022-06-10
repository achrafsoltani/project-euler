package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	file, err := os.Open("./nums.txt")
    if err != nil {
        //log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
	var sum float64 = 0
    for scanner.Scan() {
		txt := scanner.Text()
		//n, _ := strconv.ParseInt(txt[40:50], 10, 64)
		n, _ := strconv.ParseFloat(txt, 10)
		sum += float64(n)
    }
	fmt.Println(sum)
}