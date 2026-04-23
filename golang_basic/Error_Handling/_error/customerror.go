package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string)(bool, error){
	if strings.TrimSpace(input) == ""{
		return  false, errors.New("Cannot empty! please input a data")
	}else{
	return true, nil
	}
}

func main(){
	var input string
	fmt.Print("Type a data: ")
	fmt.Scanln(&input)
	if valid, err := validate(input); valid {
		fmt.Println("Your Data : ", input)
	}else{
		fmt.Println(err.Error())
	}

}