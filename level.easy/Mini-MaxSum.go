package main

import "fmt"

var valorMaxAux int = -1
var valorMinAux int = 9999999999

func main()  {

	var tamanio int = 5
	var valor int
	/*var valorMax int
	var valorMin int*/


	//fmt.Scanf("%d", &tamanio)
	arr := make([]int, tamanio)

	for i:= 0; i < tamanio; i++ {
		fmt.Scanf("%d", &valor)
		arr[i] = valor
	}

	for j := 0; j < tamanio; j++ {
		result := calcularValorArreglo(j, arr)
		validarMinValor(result)
		validarMaxValor(result)
	}

	fmt.Println(valorMinAux," ", valorMaxAux)

}

func calcularValorArreglo(index int, arr []int)  int{

	var suma int
	for i := 0; i < len(arr); i++ {

		if i != index {
			suma += arr[i]
		}

	}
	return suma
}

func validarMinValor(valorMinimo int) {

	if valorMinimo < valorMinAux {
		valorMinAux = valorMinimo
	}

}

func validarMaxValor(valorMaximo int) {

	if valorMaximo > valorMaxAux {
		valorMaxAux = valorMaximo
	}

}
