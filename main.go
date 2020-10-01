package main

import "fmt"

func main() {
	var e, fac, i, j, a float64
	e = 0
	fmt.Scan(&a)
	for i = 0; i < a; i++ {
		fac = 1
		for j = 1; j < i; j++ {
			fac = fac * j
		}
		//.Println(fac);
		e = e + (1 / fac)
	}
	fmt.Println(e)
}
