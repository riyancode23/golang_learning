package main

import "fmt"

func main() {
	var nilai = 6
	var hasil string
	if nilai == 10 {
		hasil = "Lulus Sempurna"
	} else if nilai <= 6 || nilai >= 10 {
		hasil = "Lulus"
	} else {
		hasil = "Lulus"
	}

	fmt.Printf("Hasil dengan Nilai %d adalah %s \n", nilai, hasil)
}
