package main

import "fmt"

func form(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func main() {
	var fav_num int
	fmt.Println("Enter any random number ")
	_, err := fmt.Scan(&fav_num)
	if err != nil {
		fmt.Println("Please enter a valid whole number")
	}
	res := form(fav_num)

	fmt.Printf("the sum of numbers from 0 upto %d is: %d\n", fav_num, res)
}
