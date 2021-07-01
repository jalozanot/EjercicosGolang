package main

import "fmt"

func main() {

	var tamanio int
	var valor int
	var d int
	var m int
	var suma int
	var count int
	fmt.Scanf("%d", &tamanio)

	arreglo := make([]int,tamanio)

	for i:= 0; i< tamanio; i++ {
		fmt.Scanf("%d", &valor)
		arreglo[i] = valor
	}

	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &d)

	for i:= 0; i< (tamanio-d+1); i++ {
		suma = 0
		for j:=0; j < d; j++ {
			suma = suma + arreglo[i+j]
		}

		if suma == m {
			count ++
		}

	}

	fmt.Println(count)
}
