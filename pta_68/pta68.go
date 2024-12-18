package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	arr := make([]float64, 0)
	for i := 0; i < n; i++ {
		num, _ := strconv.ParseFloat(parts[i], 64)
		arr = append(arr, num)
	}
	sum := 0.00
	for _, v := range arr {
		sum += 1 / v
	}
	result := sum / float64(n)
	result = 1 / result

	fmt.Println(result)
}
