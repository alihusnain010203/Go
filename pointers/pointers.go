package main

import "fmt"

func main() {
	// var a int = 20
	// ip := &a

	// fmt.Println("Address of a variable: ", &a)
	// fmt.Println("Address stored in ip variable: ", ip)
	// fmt.Println("Value of *ip variable: ", *ip)
	age := 32
	//    var ageP *int
	//    ageP=&age
	ageP := &age
	fmt.Println("Age :", *ageP)
	fmt.Println("Age Adress:" ,ageP)
	getAdultAge(ageP)
   	fmt.Println("Age :", *ageP)
	fmt.Println("Age :" ,age)
}

func getAdultAge(age *int)  {
	fmt.Print(age)
    *age=*age-18;
	// return *age + 18
}
