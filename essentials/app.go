package main

import (
	"fmt"
)

func main() {
	var expenses, revenue, tax int
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter tax: ")
	fmt.Scan(&tax)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax - tax

    ratio:=earningBeforeTax/earningAfterTax;
    // Display output in one line
    fmt.Printf("Earning before tax: %d, Earning after tax: %d, Ratio: %f\n", earningBeforeTax, earningAfterTax, ratio)

}
