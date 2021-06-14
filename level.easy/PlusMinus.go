package main

import "fmt"

func main()  {

	var tamanio int
	var valor int
	var countNumberNeutro int
	var countNumberPositivo int
	var countNumberNegative int

	fmt.Scanf("%d",&tamanio)
	arr := make([]int,tamanio,tamanio)

	for i := 0; i < tamanio; i++ {
		fmt.Scanf("%d", &valor)
		arr[i] = valor
	}

	for j := 0; j < tamanio; j++ {
		if arr[j] == 0 {
			countNumberNeutro ++
		} else if arr[j] < 0 {
			countNumberNegative ++
		} else if arr[j] > 0 {
			countNumberPositivo ++
		}
	}

	valorNegativo := fmt.Sprintf("%.6f",(float64(countNumberNegative) / float64(tamanio)))
	valorPositivo := fmt.Sprintf("%.6f",(float64(countNumberPositivo) / float64(tamanio)))
	valorNeutro := fmt.Sprintf("%.6f",(float64(countNumberNeutro) / float64(tamanio)))

	fmt.Println(valorPositivo)
	fmt.Println(valorNegativo)
	fmt.Println(valorNeutro)

}


