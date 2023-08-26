package main

import "fmt"

func main(){
	x := 7
	if x > 10 {
		fmt.Println("x lebih dari 10")
	} else if x > 5 {
		fmt.Println("x lebih dari 5")
	} else {
		fmt.Println("x tidak lebih dari 5")
	}
}