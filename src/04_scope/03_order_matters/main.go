package main

import (
"fmt"
)

func main() {
	
	fmt.Println(x)
	var x = 42
	fmt.Println(y)
}

var y = "Some String"


// compiles when x = 42 is moved above fmt.Println(x)