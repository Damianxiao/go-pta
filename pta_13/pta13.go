package main

func main(){
	var num int
	fmt.Scan(&num)
	for i:=1;i<n;i++{
		fmt.Println(level(i))
	}
}

func level(n int) int{
	sum := n
	for i:=n-1;i>1;i--{
		sum *= n
	}
	return sum
}