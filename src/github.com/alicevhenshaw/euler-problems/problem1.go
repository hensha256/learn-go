package main

import "fmt"

func ProblemOne(upTo int) (sum int) {
	sum = 0
	if upTo > 0 {
		for i:=1; i<upTo; i++ {
			if i%3 == 0 || i%5 == 0 {
				sum += i
			}
		}
	}
	return
}