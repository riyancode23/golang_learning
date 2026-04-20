package main

import "fmt"
import "strings"

func main(){
 var skillCoding = []string{"Php","Golang"}
 printMessage("Belajar Coding : ",skillCoding)
}

func printMessage(message string, skill []string){
 var data_skill = strings.Join(skill, ",")
 fmt.Println(message,data_skill)
}