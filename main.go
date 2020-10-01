package main

import "fmt"

func main() {
	var e, fac, i, j float64;
	e=0;
	for i = 1; i<1000; i++ {
		fac = 1
		for j = 1; j < i; j++ {
			fac = fac * j
		}
		//.Println(fac);
		e =e+ (1 / fac)
	}
	fmt.Println(e)
}
