package main

import "fmt"

func ProblemFour() (largestPalindrome int) {
	result := 0
	for i:=100; i<1000; i++ {
		for j:=100; j<1000; j++ {
			result = i*j
			if isPalindrome(result) && result > largestPalindrome {
				largestPalindrome = result
				fmt.Println(largestPalindrome)
			}
		}
	}
	return
}

func isPalindrome(number int) bool {
	temp := number
	reverse := 0
	remainder := 0
	for temp > 0 {
		remainder = temp%10
		reverse = reverse*10 + remainder
		temp = temp / 10 
	}
	return number == reverse
}