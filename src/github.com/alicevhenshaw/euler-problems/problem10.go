package main

import "fmt"

func ProblemTen(upTo int) (sum int) {
	isNotPrime := make([]bool, upTo)
	isNotPrime[0] = true
	for i:=2; i<upTo; i++ {
		if !(isNotPrime[i-1]) {
			sum += i
			for sieve:=2*i; sieve<=upTo; sieve+=i {
				isNotPrime[sieve-1] = true
			}
		}
	} 
	return
}