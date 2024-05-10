package main

import (
	"fmt"
)

// dessa vez eu fiz uma func bonitinha pra isso inves de ficar so jogando tudo na main

func main() {

	var x int 
	var y float64

	fmt.Scan(&x,&y)

	fmt.Println(consumo_Medio(x, y))
}

func consumo_Medio(x int, y float64) float64 {

	z := float64(x)
	return z/y
}
