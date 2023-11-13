package main

import "fmt"

func numDiv(number int, limit int) int {
	nDivisors := limit/number
	return (number*nDivisors*(nDivisors+1))/2
}

func main() {
	//Brute Force
	sum := 0
	for i:= 3; i < 1000; i++{
		if i % 3 == 0 || i % 5 == 0{
			sum += i
		}
	}
	fmt.Println(sum)
	//Factors Solution
	sum2 := numDiv(3, 999) + numDiv(5, 999) - numDiv(15, 999)
	fmt.Println(sum2)
}
