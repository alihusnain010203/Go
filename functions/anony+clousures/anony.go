package anpnymous

import "fmt"

func anony() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original Numbers: ", numbers)
	dNumbers := TransformNumber(&numbers, func(number int) int {
		return number * 2 	
	})

	func (){
		fmt.Println("Anon Numbers: ", dNumbers)
	}();

}

func TransformNumber(number *[]int, Transformer func(int) int) []int {
	

	dNumbers := []int{}
	for _, value := range *number {
		dNumbers = append(dNumbers, Transformer(value))
	}
	return dNumbers

}
