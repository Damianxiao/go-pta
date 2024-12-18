package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Fields(scanner.Text())
	fmt.Print(len(arr))
	if len(arr) != 6 {
		log.Fatalf("error: expected 6 elements, got %d", len(arr))
	}

	// 将字符串转换为整数
	low, err := strconv.Atoi(arr[4])
	if err != nil {
		log.Fatalf("error converting low to integer: %v", err)
	}
	dex, err := strconv.Atoi(arr[5])
	if err != nil {
		log.Fatalf("error converting dex to integer: %v", err)
	}

	// 只取前四个元素，并转换为整数
	nums := make([]int, 4)
	for i := 0; i < 4; i++ {
		nums[i], err = strconv.Atoi(arr[i])
		if err != nil {
			log.Fatalf("error converting element to integer: %v", err)
		}
	}

	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] < low {
			m[i] = true
		}
		for j := i + 1; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) > float64(dex) {
				m[i] = true
			}
		}
	}

	switch len(m) {
	case 0:
		fmt.Println("normal")
	case 1:
		for i := range m {
			fmt.Printf("warning %d\n", i)
		}
	default:
		fmt.Println("all")
	}
}
