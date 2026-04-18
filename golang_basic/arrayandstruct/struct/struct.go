package main

import "fmt"

type Spesification struct {
	merk string
	processor string
	ram int
}

func main() {
	//variable maps array
	var Smartphone Spesification
	Smartphone.merk = "Samsung"
	Smartphone.processor = "Ecynos 1330"
	Smartphone.ram = 4

	fmt.Println("Smartphone : ")
	fmt.Println("Merek \t\t:", Smartphone.merk)
	fmt.Println("Prosessor \t:", Smartphone.processor)
	fmt.Println("RAM/Memori \t:", Smartphone.ram)



	
}