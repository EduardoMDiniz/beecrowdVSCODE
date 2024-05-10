package main

import (
	"fmt"
)

func main (){

	var a int 
	var b int
	var c int 

	fmt.Scan(&a,&b,&c)

	switch {
	case a > b && a > c:
		fmt.Printf("%v eh o maior",a)

	case b > a && b > c:
		fmt.Printf("%v eh o maior",b)
	
	case c > a && c > b:
		fmt.Printf("%v eh o maior",c)
	}
}