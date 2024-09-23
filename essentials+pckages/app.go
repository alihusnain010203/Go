package main

import (
	"fmt"
)

func main() {
	// var  tax int
	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)
	expenses := getUserInfo("expenses")
	revenue := getUserInfo("revenue")
	tax := getUserInfo("Tax")

	// Display output in one line
	// fmt.Printf(`Earning before tax: %d,
	//  Earning after tax: %d,
	//   Ratio: %v`, earningBeforeTax, earningAfterTax, ratio)

	// Storing string in variable
	earningBeforeTax, earningAfterTax, ratio := performCalculation(revenue, expenses, tax)

	RatioWithString := fmt.Sprintf("Ratio is: %v", ratio)
	fmt.Printf("Earning before tax: %.2f\n, Earning after tax: %.2f\n, Ratio: %.2f\n", earningBeforeTax, earningAfterTax, ratio)
	fmt.Print(RatioWithString)
	sum := add(1, 2)
	fmt.Printf("\nSum: %d", sum)
}

func add(a int, b int) int {
	return a + b
}

func getUserInfo(text string) (inputVal float64) {
	var InputVal float64
	fmt.Print("Enter value of ", text)
	fmt.Scan(&InputVal)
	return InputVal
}
func performCalculation(revenue, expenses, tax float64) (earningBeforeTax, earningAfterTax, ratio float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax - tax

	ratio = earningBeforeTax / earningAfterTax
	return earningBeforeTax, earningAfterTax, ratio
}
