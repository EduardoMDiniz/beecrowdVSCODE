package main 

import (
	"fmt"
	"math"
)

var pi float64 =  3.14159

func main (){
	var a float64
	var b float64
	var c float64

	fmt.Scan(&a,&b,&c)

	var triangulo float64 = a*c /2                  //a = base, c = altura
	var circulo float64 = math.Pow(c,2) * pi        // c = raio
	var trapezio float64 = b * c + a / 2            // a = base1, b = base2, c = altura
	var quadrado float64 = math.Pow(b,2)            // b = lado
	var retangulo float64 = c * a                   // c = base, a = lado

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f",triangulo,circulo,trapezio,quadrado,retangulo)


}