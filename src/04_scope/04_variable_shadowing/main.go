package main

import (
"fmt"
)

func add(num int) int {
	num += 7
	return num
	
}

func main() {
	num := 10
	
	fmt.Println(add(num))
	fmt.Println(add(num))
}

