package main

import "fmt"

func main (){
var getMinMax = func(n []int) (int,int){
	var min, max int
for i ,e := range n{
	switch{
	case i == 0:
		max, min = e ,e
	case e > max:
		max= e
	case e < min:
		min = e
	}
}
return min,max
}

var numbers = []int{4,6,9,7,11,9,3}
var min, max = getMinMax(numbers)
fmt.Printf("Minimal = %d \n", min)
fmt.Printf("Maksimal = %d", max)

}