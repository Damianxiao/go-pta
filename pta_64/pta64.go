package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Trim(parts[i], " ")
		runes := []rune(parts[i])
		for j, r := range runes {
			if unicode.IsUpper(r) && r != 'I' {
				runes[j] = unicode.ToLower(r)
			}
		}

	}

}
