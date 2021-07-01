package main

import "fmt"

func main() {

	var tamanio int
	var valor int
	var valorAux int
	fmt.Scanf("%d", &tamanio)
	arreglo := make([]int, tamanio)
	arregloAux := make([]int, tamanio)


	for i:= 0; i < tamanio; i++ {
		fmt.Scanf("%d", &valor)
		arreglo[i] = valor
	}

	for i,v := range arreglo {

		valorAux = v % 5

		if (v + ( 5 - valorAux )) < 40 {
			arregloAux[i] = v
			continue
		}

		if valorAux >= 3 {
			arregloAux[i] = v + ( 5 - valorAux )
		}else if valorAux == 0 {
			arregloAux[i] = v
		}else if valorAux < 3 {
			arregloAux[i] = v
		}


	}

	for _,valor := range arregloAux {
		fmt.Println(valor)
	}


}
