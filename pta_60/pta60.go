package main

import "fmt"

func main() {
	const X int = 100
	const Y int = 100
	var x int
	var y int
	fmt.Scan(&x, &y)
	area1 := (100 - x) * (100 - y)
	area2 := x * (100 - y) / 2
	area3 := y * (100 - x) / 2
	area := X*Y - area1 - area2 - area3
	fmt.Println(area)
}
