package main

import (
"fmt"
)


const (
    _ = iota
	b = iota * 10
	c = iota * 20
)

func main() {
	fmt.Println("b" , b)
	fmt.Println("c" , c)
}
