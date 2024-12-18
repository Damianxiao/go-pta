package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	ratio := 0.0
	sign1, sign2 := 1.0, 1.0
	fmt.Scan(&num)
	if num < 0 {
		sign1 = 1.5
		num = -num
	}
	str := strconv.Itoa(num)
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '2' {
			count++
		}
	}
	ratio = float64(count) / float64(len(str)) * 100

	if num%2 == 0 {
		sign2 = 2
	}
	ratio *= sign1 * sign2
	fmt.Println("%.2f\n", ratio)
}
