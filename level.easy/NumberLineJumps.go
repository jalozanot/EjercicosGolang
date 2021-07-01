package main

import "fmt"

func main() {

	var x1 int
	var v1 int

	var x2 int
	var v2 int

	var kangarooA int
	var kangarooB int

	fmt.Scanf("%d", &x1)
	fmt.Scanf("%d", &v1)
	fmt.Scanf("%d", &x2)
	fmt.Scanf("%d", &v2)

	for i := 0; i <= 10000; i++ {
		kangarooA = (v1 * i) + x1
		kangarooB = (v2 * i) + x2

		if kangarooA == kangarooB {
			fmt.Println("YES")
			break
		}


	}

	if kangarooA != kangarooB {
		fmt.Println("NO")
	}



}