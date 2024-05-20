package main

import(
	"fmt"
)

type combustivel struct{
	receptor_metodo int
	alcool int
	gasolina int
	diesel int
}

var combusteiveis_posto combustivel = combustivel{
	receptor_metodo: 0,
	alcool: 0, 
	gasolina:  0,
	diesel:  0,
}


func main (){

	fmt.Printf("Digite o combustível que deseja: 1 Álcool 2 Gasolina 3 Diesel")
	fmt.Scan(&combusteiveis_posto.receptor_metodo)

	combusteiveis_posto.enquete(combusteiveis_posto.receptor_metodo)

	fmt.Printf("MUITO OBRIGADO\nAlcool: %v\nGasolina: %v\nDiesel: %v", combusteiveis_posto.alcool,combusteiveis_posto.gasolina,combusteiveis_posto.diesel)

}

func(x *combustivel) enquete (tipo int) {

	switch {
		case tipo == 1:
			(*x).alcool++
		case tipo == 2: 
			(*x).gasolina++
		case tipo == 3:
			(*x).diesel++
		default:
			fmt.Println("Código Inválido")
			fmt.Println("Por favor, insira um novo código:")
			fmt.Scan(&combusteiveis_posto.receptor_metodo)
			combusteiveis_posto.enquete(combusteiveis_posto.receptor_metodo)
	}
}
