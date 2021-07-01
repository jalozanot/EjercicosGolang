package main

import "fmt"

func main() {

	var valor int
	var s int
	var t int
	var a int
	var b int
	var arrayApplesTam int
	var arrayOrangesTam	int


	fmt.Scanf("%d", &s)
	fmt.Scanf("%d", &t)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &arrayApplesTam)
	fmt.Scanf("%d", &arrayOrangesTam)

	arregloApple := make([]int, arrayApplesTam)
	arregloOrange := make([]int, arrayOrangesTam)
	var countA int
	var countB int

	for i:= 0; i < arrayApplesTam; i++ {
		fmt.Scanf("%d", &valor)
		arregloApple[i] = valor
	}

	for i:= 0; i < arrayOrangesTam; i++ {
		fmt.Scanf("%d", &valor)
		arregloOrange[i] = valor
	}

	for _, valor := range arregloApple {
		if valor+a >= s && valor+a <= t {
			countA ++
		}
	}

	for _, valor := range arregloApple {
		if valor+b >= s && valor+b <= t {
			countB ++
		}
	}


	fmt.Println(countA)
	fmt.Println(countB)


}