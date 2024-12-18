package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const NUM1 float64 = 2.455
	const NUM2 float64 = 1.266
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	arr := make([]float64, 0)
	for _, v := range parts {
		temp, _ := strconv.ParseFloat(v, 64)
		arr = append(arr, temp)
	}
	r := 0.00
	if arr[1] == 0 {
		r = arr[0] * NUM1
	} else if arr[1] == 1 {
		r = arr[0] * NUM2
	}
	if r > arr[2] {
		fmt.Println("%f,T", r)
	} else {
		fmt.Println("%f,V", r)
	}
}
