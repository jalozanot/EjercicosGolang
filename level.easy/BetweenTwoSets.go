package main

import "fmt"

func main() {

	var lengthArrA int
	var lengthArrB int
	var valor int
	var count int = 0
	var flag bool = true
	fmt.Scanf("%d",&lengthArrA)
	fmt.Scanf("%d",&lengthArrB)

	arregloA := make([]int,lengthArrA)
	arregloB := make([]int,lengthArrB)


	for i:= 0; i < lengthArrA; i++ {
		fmt.Scanf("%d", &valor)
		arregloA[i] = valor
	}

	for i:= 0; i < lengthArrB; i++ {
		fmt.Scanf("%d", &valor)
		arregloB[i] = valor
	}


	for i:= 1; i <= 100; i++ {
		flag = true
		for _,valor := range arregloA{

			if i % valor != 0 {
				flag = false
				break
			}

		}

		if flag == true {
			for _,valor := range arregloB{

				if  valor % i != 0 {
					flag = false
					break
				}

			}
		}

		if flag == true {
			count ++
		}

	}

	fmt.Println(count)

}
