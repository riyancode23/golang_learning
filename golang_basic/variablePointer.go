package main

import "fmt"

func main(){
	// Variable Pointer
	var nilaiA int = 15
	var nilaiB *int = &nilaiA

	fmt.Println("Jika diambil data alamat Memori", nilaiB)
	fmt.Println("Jika diambil isi data variable", *nilaiB)


}