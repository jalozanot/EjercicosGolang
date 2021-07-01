package main

import "fmt"

func main() {

	var tamanio int
	var valor int
	var count int
	fmt.Scanf("%d", &tamanio)

	var k int
	fmt.Scanf("%d", &k)

	arreglo := make([]int,tamanio)

	for i:=0 ; i < tamanio; i++ {
		fmt.Scanf("%d", &valor)
		arreglo[i] = valor
	}


	for i, valorI := range arreglo {

		for j, valorJ := range arreglo {

			if i != j {

				if i < j {
					if (valorI + valorJ) % k == 0 {
						count ++
					}
				}

			}
		}

	}

	fmt.Println(count)

}
