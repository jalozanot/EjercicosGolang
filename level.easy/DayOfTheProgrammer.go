package main

import "fmt"

func main() {

	var fecha int
	fmt.Scanf("%d",&fecha)


	if fecha >= 1700 && fecha <= 1917 {

		if (fecha % 4) == 0 {
			fmt.Printf("12.09.%d",fecha)
		} else {
			fmt.Printf("13.09.%d",fecha)
		}

	} else if fecha == 1918{
		fmt.Println("26.09.1918")
	}else {

		if fecha % 400 == 0 {
			fmt.Printf("12.09.%d",fecha)
		}else if fecha % 4 == 0 && fecha % 100 != 0 {
			fmt.Printf("12.09.%d",fecha)
		} else {
			fmt.Printf("13.09.%d",fecha)
		}

	}



}
