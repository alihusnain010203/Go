package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1,2,3,4,5)
	println(sum)


}

// func sumup(numbers *[]int) int {
// 	sum := 0
// 	fmt.Println((*numbers)[0])
// 	for _, number := range *numbers {
// 		sum += number
// 	}
// 	return sum
// 	// return numbers[0] + sumup(numbers[1:])
// }

// variadic functions
func sumup(numbers ...int) int {
	fmt.Println(numbers)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}