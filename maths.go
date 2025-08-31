package main

import "fmt"

// Function to add two integers
func addIntegers(a, b int) int {
	return a + b
}

// Function to subtract two integers
func subtractIntegers(a, b int) int {
	return a - b
}

// Function to multiply two float64 numbers
func multiplyFloats(a, b float64) float64 {
	return a * b
}

// Function to divide two float64 numbers (returns an error if dividing by zero)
func divideFloats(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main1() {
	// Integer variables
	num1 := 10
	num2 := 5

	// Float variables
	float1 := 12.5
	float2 := 2.5

	// Perform operations
	sum := addIntegers(num1, num2)
	difference := subtractIntegers(num1, num2)
	product := multiplyFloats(float1, float2)
	quotient, err := divideFloats(float1, float2)

	// Print results
	fmt.Printf("Addition (%d + %d) = %d\n", num1, num2, sum)
	fmt.Printf("Subtraction (%d - %d) = %d\n", num1, num2, difference)
	fmt.Printf("Multiplication (%.2f * %.2f) = %.2f\n", float1, float2, product)

	// Handle division error (if any)
	if err != nil {
		fmt.Println("Division Error:", err)
	} else {
		fmt.Printf("Division (%.2f / %.2f) = %.2f\n", float1, float2, quotient)
	}
}
