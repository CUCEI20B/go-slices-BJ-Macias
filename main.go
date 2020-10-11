package main

import "fmt"

func main()  {

	var s []int
	var aux int
	fmt.Scan(&aux)
	var aux2 int
	var suma int
	for  i := 0; i < aux; i++ {
		fmt.Scan(&aux2)
		s = append(s, aux2)
		suma += aux2
	}
	println(suma)
}