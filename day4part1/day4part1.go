package main

import (
	"fmt"
	"strconv"
)

const min = 206938
const max = 679128

func main() {
	/* Test
	fmt.Println(isValid(111111))
	fmt.Println(isValid(223450))
	fmt.Println(isValid(123789))
	*/
	count := 0
	for n := min; n <= max; n++ {
		if isValid(n) {
			count++
		}
	}
	fmt.Println(count)
}

func isValid(p int) bool {
	password := strconv.Itoa(p)
	dec := true
	repeat := false
	for i := 1; i < len(password); i++ {
		if password[i] < password[i-1] {
			dec = false
			break
		}
		if !repeat && password[i] == password[i-1] {
			repeat = true
		}
	}
	return (dec && repeat)
}
