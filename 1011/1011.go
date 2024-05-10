package main

import (
	"fmt"
	"math"
)

var pi float64 = 3.14159

func main (){

	var raio float64 

	fmt.Scan(&raio)

	var volume float64 = 4.0/3.0 * pi * math.Pow(raio,3)

	fmt.Printf("VOLUME = %.3f",volume)


}
