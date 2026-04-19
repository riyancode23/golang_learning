package main

import "fmt"

func main() {
	//variable maps array
	var omset map[string]int
	 omset = map[string]int{}
	 omset["January"] = 100
	 omset["Desember"] = 500
	 fmt.Println("Omset Penjualan Januari = ", omset["January"])

}