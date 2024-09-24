package main

import (
	"fmt"
	"example.com/structs/user"
)


func main() {
	userfirstName := getUserInput("Enter Your Firstname.!!!")
	userlastName := getUserInput("Enter Your Lastname!!!")
	userbirthday := getUserInput("Enter Your Date Of Birth DD/MM/YYYY!!!")

	var Appuser *user.User
	// Appuser = User{
	// 	firstName: userfirstName,
	// 	lastName:  userlastName,
	// 	birthday:  userbirthday,
	// 	createdAt: time.Now(),
	// }

	Appuser,err :=user.New(userfirstName,userlastName,userbirthday)
	if(err!=nil){
		fmt.Println(err)
		panic(err)
	}
	
	// Output(&Appuser)

	Appuser.OutputStruct()
	Appuser.Clear()
	Appuser.OutputStruct()

}

// get Value of Struct and the copy og user created in memory
//
//	func Output(user User) {
//		fmt.Println(user.firstName, user.lastName, user.birthday,user.createdAt)
//	}
// func Output(user *User) {
// 	fmt.Println(user.firstName, user.lastName, user.birthday, user.createdAt)
// 	// use de refernce
// 	fmt.Println((*user).firstName, (*user).lastName, (*user).birthday, (*user).createdAt)

// }
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	return value

}
