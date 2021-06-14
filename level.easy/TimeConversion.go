package main

import (
	"fmt"
	"time"
)

func main() {


	var fechaString string
	fmt.Scanf("%s",&fechaString)

	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, _ := time.Parse(layout1, fechaString)


	fmt.Println(t.Format(layout2))


}
