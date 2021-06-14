package main

import "fmt"

func main()  {

	var tamanio int
	var valor int
	fmt.Scanf("%d",&tamanio)

	arr := make([]int, tamanio)

	for i:=0; i < tamanio; i++ {
		fmt.Scanf("%d",&valor)
		arr[i] = valor
	}

	numMayor := encontrarNumMayor(arr)

	contadorNumMayor := contarNumMayor(numMayor, arr)

	fmt.Println(contadorNumMayor)

}


func encontrarNumMayor(arreglo []int) int{


	var numMayor = -1

	for _,v := range arreglo {

		if v > numMayor {
			numMayor = v
		}

	}

	return numMayor
}

func contarNumMayor(numMayor int, arr []int) int{

	var countNumMayor int
	for _, v := range arr {
		if v == numMayor{
			countNumMayor ++
		}
	}
return countNumMayor
}