package main

import (
"fmt"
)

var x = 0

func increment() {
	x++
	fmt.Println(x)
}
func main () {
	increment()
	increment()
}
