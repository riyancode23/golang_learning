package main

import "fmt"

func main () {
	var productMerk string = "Asus"
	var price uint32 = 0

	switch productMerk {
	case "Samsung":
		price = 24000000
	case "Xiaomi":
		price = 15000000
	case "Asus":
		price = 17500000
	default:
		price = 0
	}

	fmt.Printf("Product Merk %s Harga %d \n", productMerk, price)

}
