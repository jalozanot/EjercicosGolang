package main

import "fmt"

func main() {

	var arregloBob[3] int
	var arregloAlice[3] int
	var valor int
	var countAlice int
	var countBob int


	for i:= 0; i < 3; i++ {
		fmt.Scanf("%d", &valor)
		arregloAlice[i] = valor
	}

	for j:= 0; j < 3; j++ {
		fmt.Scanf("%d", &valor)
		arregloBob[j] = valor
	}

	for i := 0; i <3 ; i++ {

		if arregloAlice[i] > arregloBob[i] {
			countAlice++
		}

		if arregloAlice[i] < arregloBob[i] {
			countBob++
		}

	}

	fmt.Println(countAlice, countBob )


}
