package funcAsValue

import (
	"fmt"
)

func funcAsValu(){
	numbers:=[]int{0,1,2,3,4,5,6,7,8,9};
	fmt.Println("Original Numbers: ",numbers);
	dNumbers:=TransformNumber(&numbers,Transformer);
	fmt.Println("Doubled Numbers: ",dNumbers);
	fmt.Println("Original Numbers: ",numbers);




}

func TransformNumber(number *[]int, Transformer func(int) func(int) int) []int {
	fmt.Println("Entered TransformNumber");
	var transformerNumber int;
	fmt.Println("Enter 1 for Doubled and 2 for Tripled");
	fmt.Scan(&transformerNumber);
	transformer:=Transformer(transformerNumber);

	dNumbers := []int{}
	for _, value := range *number {
		dNumbers = append(dNumbers, transformer(value))
	}
	return dNumbers;

}

func Transformer(value int) func(int) int{
	return func(number int) int{
		return number*value;
	}
}

func Doubled(number int) int{
	return number*2;
} 	 

func Tripled(number int) int{
	return number*3;
}