package main

import (
	"fmt"
)

func main() {

	var funcionario int
	var horas int
	var salarioHoraTrabalhada float64

	fmt.Scan(&funcionario, &horas, &salarioHoraTrabalhada)
    
    horasFloat := float64 (horas)
	var salario float64 = horasFloat * 10.00
	fmt.Printf("NUMBER = %v\nSALARY = U$ %v\n", funcionario,salario)
}
