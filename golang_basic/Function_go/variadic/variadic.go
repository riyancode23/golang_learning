package main

import "fmt"

func main (){
 var avg = calculate(10,22,31,77,55,15)
 var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
 fmt.Println(msg)
}

//function variadic adalah fungsi yang menggunakan parameter tak terbatas(array)
func calculate(numbers ...int) float64 {
var total int = 0
for _, numbers := range numbers {
	total+=numbers
}

var avg = float64(total) / float64(len(numbers))

return  avg
}