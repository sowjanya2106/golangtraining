package main

import "fmt"

func main () {
	
	a := "hi"
	
	switch  "hi" {
		case len(a) == 2 :
			fmt.Println("hi !!")
		case a == "hello" :
			fmt.Println("hello !!")
		case a == "hola":
			fmt.Println("hola !!")	
		default 
			fmt.Println("bye !!")		
	}
}