package main

import "fmt"
import "math"

func main(){
 var diameter float32= 4
 luas, keliling := calculate_circle(diameter)
 fmt.Printf("Luas Lingkaran D=%f adalah %f ",diameter,luas)
 fmt.Printf("Keliling Lingkaran D=%f adalah %f ",diameter,keliling)
}

func calculate_circle(diameter float64)(float64,float64){
 var luas = math.Pi * math.Pow(diameter/2, 2) // luas lingkaran
 var keliling = math.Pi * diameter // keliling lingkaran
 
 return luas , keliling
} 