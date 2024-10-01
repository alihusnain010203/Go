package lists

import (
	"fmt"
)

// func main() {
// 	// var prices [4]float64
// 	// prices[0] = 100.5
// 	// prices[1] = 200.0
// 	// prices[2] = 300.0
// 	// prices[3] = 400.0
// 	prices := [4]float64 {100.5, 200.0, 300.0}

// 	// slices
// 	// featuredPrices := prices[1:3]
// 	// omitLast := prices[0:4]
// 	// omitFirst := prices[1:]
// 	// fmt.Println(omitFirst)
// 	// fmt.Println(omitLast)
// 	// fmt.Println(featuredPrices)
// 	// fmt.Println(prices[3])
// 	featuredPrices := prices[2:3]
// 	featuredPrices = append(featuredPrices, 500.0)
//     featuredPrices[0] = 500.0
// 	fmt.Println(featuredPrices)
// 	fmt.Println(prices)

// 	length := len(featuredPrices)
// 	fmt.Println(length)
// 	capacity := cap(featuredPrices)
// 	fmt.Println(capacity)
// }

// func main(){
// 	// dynamic array
// 	prices := []float64{100.5, 200.0, 300.0}
// 	prices = append(prices, 400.0)
// 	fmt.Println(prices)
// }

func main(){
	// 1)
	hobbies:=[3]string{"reading", "cooking", "coding"}
	// fmt.Println(hobbies)
	// 2)
	// fmt.Println(hobbies[0])
	// 3)
	// featuredHobbies := [2]string{hobbies[1], hobbies[2]}
	// fmt.Println(featuredHobbies)
	// 4)
	// slicesHobbies := hobbies[0:2]

	slicesHobbies:=hobbies[:2]
	// slicesHobbies = slicesHobbies[2:3]
	// slicesHobbies=slicesHobbies[1:3]/


	fmt.Println(slicesHobbies)
}