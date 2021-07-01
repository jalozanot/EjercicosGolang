package main

import "fmt"

var countTipoUno int
var countTipoDos int
var countTipoTres int
var countTipoCuatro int
var countTipoCinco int

var typeBird map[int]int
var numMayor int = -1
func main() {

	var tamArreglo int
	var valor int
	var numMenor int


	fmt.Scanf("%d", &tamArreglo)

	typeBird = make(map[int]int)
	for i:=0 ; i < tamArreglo; i++ {
		fmt.Scanf("%d", &valor)
			switch valor {
			case 1:
				countTipoUno ++
				typeBird[1] = countTipoUno
			case 2:
				countTipoDos ++
				typeBird[2] = countTipoDos
			case 3:
				countTipoTres ++
				typeBird[3] = countTipoTres
			case 4:
				countTipoCuatro ++
				typeBird[4] = countTipoCuatro
			case 5:
				countTipoCinco ++
				typeBird[5] = countTipoCinco
			}
	}

	for i:=1 ; i <= 5 ; i++{

		if typeBird[i] > numMayor {
			numMayor = typeBird[i]
			numMenor = i
		}

	}
	fmt.Println(numMenor)
}


