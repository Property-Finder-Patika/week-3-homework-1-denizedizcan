package main

import "fmt"

func sieve(n int) {
	var prime []bool
    for i:=0; i<n+1; i++ {
        prime = append(prime, true)
    }
	p := 2
	
	for p*p <= n{
		if prime[p] == true{

			for i:= p*p; i< n+1; i+=p{
				prime[i] = false
			}
		
		}
		p += 1
	}
	for i := 2; i< n+1; i++{
		if prime[i]{
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
    nums := 20

    sieve(nums)

}