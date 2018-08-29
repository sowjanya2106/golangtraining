package main

import (
"fmt"
)


const ( 
	a = iota
	b = iota * 10
	c = iota * 20
)

const ( 
	d = iota
	e
	f
)

func main () {
	fmt.Println("a" ,a)
	fmt.Println("b" ,b)
	fmt.Println("c" ,c)
	fmt.Println("d" ,d)
	fmt.Println("e" ,e)
	fmt.Println("f" ,f)
}
