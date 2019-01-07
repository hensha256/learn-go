package main

import "fmt"

func ProblemSeven(upTo int) int {
	primesFound := 0
	nextGuess := 1
	for (primesFound < upTo) {
		nextGuess += 1
		if nextGuess>2 && nextGuess%2==0 {
			nextGuess++
		}
		if isPrime(nextGuess) {
			primesFound++
		}
	}
	return nextGuess
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	} else {
		for i:=2; i<number; i++ {
			if number%i==0 {
				return false
			}
		}
	}
	return true
}