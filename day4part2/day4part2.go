package main

import (
	"fmt"
	"strconv"
)

const min = 206938
const max = 679128

func main() {
	/* Test
	fmt.Println(isValid(112233))
	fmt.Println(isValid(123444))
	fmt.Println(isValid(111122))
	fmt.Println(isValid(112222))
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
	count := 0
	for i := 1; i < len(password); i++ {
		if password[i] < password[i-1] {
			dec = false
			break
		}
		if password[i] == password[i-1] {
			count++
		} else if count == 1 {
			repeat = true
		} else {
			count = 0
		}
	}
	// End Case
	if count == 1 {
		repeat = true
	}
	return (dec && repeat)
}
