package main

import "fmt"

func main () {
	switch  "hello" {
		case "hi" , "hello", "hola" :
			fmt.Println("Greetings")
		default 
			fmt.Println("bye !!")		
	}
}