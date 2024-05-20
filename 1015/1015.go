package main 

import (
	"fmt"
	"math"
)

var x1 float64
var y1 float64
var x2 float64
var y2 float64

func main (){

	fmt.Printf("Insira os valores de P1 (x1 e y1)")
	fmt.Scan(&x1, &y1)
	fmt.Printf("Insira os valores de P2 (x2 e y2)")
	fmt.Scan(&x2, &y2)

	distancia_entre_eles(x1,y1,x2,y2)
}

func distancia_entre_eles (pontos...float64) {

	distancia0 := math.Pow((pontos[1]) - pontos[0], 2) + math.Pow(pontos[3] - pontos[2], 2)
	distancia := math.Sqrt(distancia0)
	
	fmt.Printf("%.2f",distancia)
}