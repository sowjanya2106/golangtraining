package main

import "fmt"

func main () {
	switch "hi" {
		case "hi" :
			fmt.Println("hi !!")
		case "hello" :
			fmt.Println("hello !!")
		case "hola" :
			fmt.Println("hola !!")	
		default 
			fmt.Println("bye !!")		
	}
}