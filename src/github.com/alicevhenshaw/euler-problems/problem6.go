package main

import "fmt"

func ProblemSix(upTo int) int {
	squaredSum := 0
	sum := 0
	sumOfSquares := 0
	for i:=1; i<=upTo; i++ {
		sum += i
		sumOfSquares += (i*i)
	}
	squaredSum = sum*sum
	return squaredSum - sumOfSquares
}