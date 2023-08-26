package main

import "fmt"

func main() {
	// Membuat map memetakan buah dan harganya
	fruitPrices := map[string]int{
		"apple": 50,
		"banana": 30,
		"orange": 40,
	}
	
	// Menampilkan harga buah-buahan
    fmt.Println("Price of apple:", fruitPrices["apple"])
    fmt.Println("Price of banana:", fruitPrices["banana"])
    fmt.Println("Price of orange:", fruitPrices["orange"])

	// Menambahkan nilai baru
	fruitPrices["grape"] = 60
	
	// Menghapus nilai dari map
	delete(fruitPrices, "banana")
	
	// Memeriksa apakah sebuah kunci ada dalam map
    if price, ok := fruitPrices["apple"]; ok {
        fmt.Println("Price of apple:", price)
    } else {
        fmt.Println("Apple not found in the map")
    }
}