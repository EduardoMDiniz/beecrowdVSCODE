package main     

import (
	"fmt"
)

func main (){

	var vendedor string 
	var salarioFixo float64
	var vendasTotais float64

	fmt.Scan(&vendedor,&salarioFixo,&vendasTotais)

	var salarioDoMes float64 = salarioFixo + vendasTotais * 0.15

	fmt.Printf("TOTAL = R$ %.2f", salarioDoMes)

}
