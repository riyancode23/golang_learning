package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string)(bool, error){
	if strings.TrimSpace(input) == ""{
		return  false, errors.New("Input cannot empty! please input a data")
	}else{
	return true, nil
	}
}

func catch(){
	if r := recover() ; r != nil{
		fmt.Println("Error occured ::", r)
	}else{
		fmt.Println("The Aplication Running Perfectly")
	}
}

func main(){
	defer catch()
	var input string
	fmt.Print("Type a data: ")
	fmt.Scanln(&input)
	if valid, err := validate(input); valid {
		fmt.Println("Your Data : ", input)
	}else{
		panic(err.Error())
		//kode dibawah kenika ada syntax golang panic maka tidak akan dijalankan/berhenti
		fmt.Println("tidak terbaca")
	}

}