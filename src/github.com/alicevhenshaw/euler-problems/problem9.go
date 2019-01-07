package main

import "fmt"

func ProblemNine() int {
	for i:=1; i<1000; i++ {
		for j:=i+1; j<1000; j++ {
			diff := 1000 - (i+j)
			squareTotal := i*i + j*j
			squareDiff := diff*diff
			if squareDiff==squareTotal {
				return i*j*diff
			}
		}
	}
	return 0
}