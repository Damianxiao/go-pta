package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入一句英文句子：")
	if scanner.Scan() {
		input := scanner.Text()

		// Step 1: 输出原句
		fmt.Println(input)

		// Step 2: 消除多余空格，行首尾空格，标点符号前空格
		processed := cleanSpaces(input)

		// Step 3: 把大写字母变小写，保留 I
		processed = convertCase(processed)

		// Step 4: 替换独立的 "can you" 和 "could you"
		processed = replaceCanCould(processed)

		// Step 5: 替换独立的 "I" 和 "me" 为 "you"
		processed = replaceIAndMe(processed)

		// Step 6: 替换问号为感叹号
		processed = strings.ReplaceAll(processed, "?", "!")

		// Step 7: 输出替换后的句子
		fmt.Println(processed)
	}
}

// 消除多余空格和标点符号前的空格
func cleanSpaces(input string) string {
	// 去除行首行尾的空格
	input = strings.TrimSpace(input)

	// 把多个空格替换为一个空格
	re := regexp.MustCompile(`\s+`)
	input = re.ReplaceAllString(input, " ")

	// 把标点符号前的空格去掉
	re = regexp.MustCompile(`\s([.,!?])`)
	input = re.ReplaceAllString(input, "${1}")

	return input
}

// 转换大小写，保留 I 的大写
func convertCase(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if word != "I" { // 保留 I
			words[i] = strings.ToLower(word)
		}
	}
	return strings.Join(words, " ")
}

// 替换独立的 "can you" 和 "could you"
func replaceCanCould(input string) string {
	re := regexp.MustCompile(`\bcan you\b`)
	input = re.ReplaceAllString(input, "I can")

	re = regexp.MustCompile(`\bcould you\b`)
	input = re.ReplaceAllString(input, "I could")

	return input
}

// 替换独立的 "I" 和 "me" 为 "you"
func replaceIAndMe(input string) string {
	re := regexp.MustCompile(`\bI\b`)
	input = re.ReplaceAllString(input, "you")

	re = regexp.MustCompile(`\bme\b`)
	input = re.ReplaceAllString(input, "you")

	return input
}
