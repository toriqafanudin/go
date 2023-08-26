package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("angka:", i)
	}
	
	k := 0
	for k < 5 {
		fmt.Println("Angka", k)
		k++
	}
	
	numbers := []int{10, 20, 30, 40, 50}
	
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}