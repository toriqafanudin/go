package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
	
	var age int = 25
	var weight float64 = 65.5
	var grade rune = 'A'
	var isPassed bool = true
	var name string = "Faisa Nirbita"
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Weight:", weight)
	fmt.Println("Grade:", grade)
	fmt.Println("Passed:", isPassed)
	
	// Membuat slice bilangan bulat
	numbers := []int{10, 20, 30, 40, 50}
	
	// Mengakses element slice dan mencetaknya
	fmt.Println("Element 1:", numbers[0])
	fmt.Println("Element 2:", numbers[1])
	fmt.Println("Element 3:", numbers[2])
	fmt.Println("Element 4:", numbers[3])
	fmt.Println("Element 5:", numbers[4])
	fmt.Println("Semua Element:", numbers)
	
}