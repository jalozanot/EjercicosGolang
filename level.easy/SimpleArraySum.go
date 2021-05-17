package main

import "fmt"

func main() {

	var tam int
	var valor int
	fmt.Scanf("%d",&tam)

	arreglo := make([]int, tam)

	for i := 0; i < tam; i++ {
		fmt.Scanf("%d", &valor)
		arreglo[i] = valor
	}

	respuesta(arreglo)

}

func respuesta(valor[] int)  {
	suma := 0

	for _,v := range valor{
		suma = suma + v
	}

	fmt.Println(suma)
}