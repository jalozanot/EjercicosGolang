package main

import (
	"fmt"
	"math"
)

func main() {

	var tam int
	var valor int
	fmt.Scanf("%d", &tam)
	 arreglo := make([][]int, tam)
	var diagonalA int
	var diagonalB int


	for i := 0; i < tam; i++ {
		arreglo[i] = make([]int, tam)
		for j := 0; j < tam; j++ {
			fmt.Scanf("%d", &valor)
			arreglo[i][j] = valor
		}
	}

	for i := 0; i < tam; i++ {
		for j := 0; j < tam; j++ {

			if i == j {
				diagonalA = diagonalA + arreglo[i][j]
			}

			if j == (tam-1-i) {
				diagonalB = diagonalB + arreglo[i][j]
			}

		}
	}

	fmt.Println(math.Abs(float64(diagonalA) - float64(diagonalB)))


}
