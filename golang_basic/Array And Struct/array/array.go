package main

import "fmt"

func main() {
	//variable array
	var coding = [4]string{"PHP","Golang","Javascript","Java"}

	fmt.Println("Jumlah Data \t\t", len(coding))
	fmt.Println("semua data \t", coding)
	fmt.Println("Ambil Data ke 0 \t", coding[0])
	fmt.Println("Ambil Data ke 1 \t", coding[1])
	fmt.Println("Ambil Data ke 2 \t", coding[2])
	fmt.Println("Ambil Data ke 3 \t", coding[3])
}