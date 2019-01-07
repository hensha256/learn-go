package main

import "fmt"

func ProblemTwo(upTo int) (sum int) {
	sum = 0
	first := 1
	second := 1
	tempFirst := 0
		for second <= upTo {
			if second%2 == 0 {
				sum += second
			}
			tempFirst = first
			first = second
			second += tempFirst
		}
	return
}