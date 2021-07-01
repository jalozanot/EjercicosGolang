package main

import "fmt"

var scoreMin int
var scoreMax int

func main()  {

	var tamanio int
	var valor int
	var countMin int
	var countMax int



	fmt.Scanf("%d", &tamanio)

	arreglo := make([]int,tamanio)

	for i:= 0; i < tamanio; i++ {
		fmt.Scanf("%d", &valor)
		arreglo[i] = valor
	}

	for index, valor := range arreglo {

		if index == 0 {
			countMin = 0
			scoreMin = valor
			scoreMax = valor
			countMax = 0

		} else {

			countMin = countMin + numberMin(scoreMin, valor)
			countMax = countMax + numberMay(scoreMax, valor)

		}

	}
	fmt.Println(countMax, countMin)

}


func numberMin(numActual, numNew int) int {

	if numNew < numActual {
		scoreMin = numNew
		return 1

	} else {
		return 0
	}

}


func numberMay(numActual, numNew int) int {

	if numNew > numActual {
		scoreMax = numNew
		return 1
	} else {
		return 0
	}

}
