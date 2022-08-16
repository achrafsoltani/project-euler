package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	fmt.Println("hello")
	x := 65 ^ 42
	fmt.Println(x)
	text := "hello world!"
	cycle := cycle_password(text, "abc")
	fmt.Println(text)
	fmt.Println(cycle)
	encoded := readFile()

	split := strings.Split(encoded, ",")
	fmt.Println("string size :", len(split))
	for i := 97; i <= 122; i++{
		for j := 97; j <= 122; j++{
			for k := 97; k <= 122; k++{
				//vars c1, c2, c3 = rune(i), rune(j), rune(k)
				//password := string(c1) + string(c2) + string(c3)
				var l int = 0
				pass_slice := []int{i, j, k}
				res_slice := []string{}
				for _, v := range split {
					element, _ := strconv.Atoi(v)
					var key int = pass_slice[l]
					var res int = element ^ key
					res_slice = append(res_slice, fmt.Sprintf("%c", rune(res)))
					//fmt.Printf("%c\n", rune(res))
					if (l == 2) {
						l = 0
					} else {
						l++
					}
				}
				joined := strings.Join(res_slice, "")
				if (!strings.Contains(joined, "}") && !strings.Contains(joined, "{") && !strings.Contains(joined, "<") && !strings.Contains(joined, ">")){
					fmt.Println(i,j,k)
					fmt.Println(joined)
					fmt.Println("#######")
				}
			}
		}
	}
}

func cycle_password(text string, password string) string{
	var text_len int = len(text)
	//var pass_len int = len(password)
	var pass_array []rune = []rune(password)
	var matching_string = ""
	var j int = 0
	for i := 0; i<text_len; i++ {
		str := pass_array[j]
		matching_string += string(str)
		if (j == 2) {
			j = 0
		} else {
			j++
		}
		
	}
	return matching_string
}

func readFile() string {
	dat, _ := os.ReadFile("./p059_cipher.txt")
    //check(err)
    return string(dat)
}