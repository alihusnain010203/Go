package main

import (
	"example.com/cal/prices"
	"fmt"
)

func main() {
	prices := []float64{10, 20, 30, 40, 50}
	taxRate := []float64{0.2, 0.3, 0.4, 0.5, 0.6}
	result := make(map[float64][]float64)

	for _, taxRate := range taxRate {
		var taxIncurred []float64 = make([]float64, len(prices))
		for priceIndex, price := range prices {

			taxIncurred[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncurred
	}

	fmt.Println(result)
	// prices:=make([]float64,1,10)
	// prices[0]=10
	// prices=append(prices,20)
	// fmt.Println(prices)
}
