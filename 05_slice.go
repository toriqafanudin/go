package main

import "fmt"

func main() {
	// Membuat slice bilangan bulat
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	
	// Menambahkan element baru
	numbers = append(numbers, 6)
	fmt.Println(numbers)
	
	// Menghapus element ke-3
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)
}