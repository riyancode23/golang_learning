package main

import "fmt"

func main() {
	//variable make array
	var city = make([]string, 3)
	city[0] = "Surabaya"
	city[1] = "Jakarta"
	city[2] = "Bandung"

	fmt.Println("List Kota : ", city)
	
}