package main

import (
	"fmt"
	"os"
)

func main(){
	defer fmt.Println("Hallo Golang")
	os.Exit(1)
	fmt.Println("Learning Now")
}