package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)
	frdm := make(map[string]bool, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		k := len(data) - 1
		for k > 1 {
			for _, id := range data[1:] {
				frdm[id] = true
			}
		}
	}
}
