package main

import (
	"fmt"
	"math"
)

func primeFactors(n int) {
    for n % 2 == 0{
		fmt.Printf("2 ")
		n = n /2
	}
	for i:= 3; i <= int(math.Sqrt(float64(n))); i+=2{
		for n%i ==0{
			fmt.Printf("%d ", i)
			n = n/i
		}
	}
	if n > 2{
		fmt.Printf("%d ", n)
	}
}

func main() {
   n := 315
   primeFactors(n)

}