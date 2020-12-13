package main

import "fmt"

func main() {

	var tam int
	var valor int64
	var suma int64
	fmt.Scanf("%d",&tam)

	arreglo := make([]int64, tam)

	for i := 0; i < tam; i++ {

		fmt.Scanf("%d",&valor)
		arreglo[i] = valor
	}
fmt.Println(len(arreglo))
	for _, v:= range arreglo {
		suma = suma + v
	}

	fmt.Println(suma)

}
