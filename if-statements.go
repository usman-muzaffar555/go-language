package main

import (
	"fmt"
)

func iffunctions() {

	// simple if statement
	fmt.Println("Basic if statement - [x=10] check if x is greater than 5, where x value is 10")
	x := 10
	if x > 5 {
		fmt.Println("result is: x is greater than 5")
	}

	// if else statement
	fmt.Println("result is: Basic if-else - check x is od or even, where x=10")
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// if-else-if statement
	fmt.Println("result is: Basic if-else-if - check if x<0 then nagative, if x=0 zero, else positive")
	if x < 0 {
		fmt.Println("Negative")
	} else if x == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive")
	}

}
