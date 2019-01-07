package main

import "fmt"

func ProblemFive() (guess int) {
	primeFactors := 2*3*5*7*11*13*17*19
	guess = 0
	found := false
	for !found {
		guess += primeFactors
		for i:=1; i<=20; i++ {
			if guess%i != 0 {
				break
			}
			if i==20 {
				found = true
			}
		}
	}
	return
}

func main() {
	fmt.Println(ProblemFive())
}