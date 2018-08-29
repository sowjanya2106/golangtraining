package main

import (
"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func main () {
	fmt.Println("a" ,a)
	fmt.Println("b" ,b)
	fmt.Println("c" ,c)
}

