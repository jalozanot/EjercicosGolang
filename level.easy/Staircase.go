package main

import "fmt"

func main() {

	var tamanio int
	fmt.Scanf("%d", &tamanio)

	for i:=1; i <= tamanio; i++ {
		for j := 1; j <= tamanio; j++ {

			if j <= (tamanio-i) {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}

		}
		fmt.Println()
	}
}