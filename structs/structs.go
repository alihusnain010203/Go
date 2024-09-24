package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}
func (user *User) OutputStruct() {
	fmt.Println(user.firstName, user.lastName, user.birthday,user.createdAt)
	// use de refernce
	fmt.Println((*user).firstName, (*user).lastName, (*user).birthday,(*user).createdAt)

}

func main() {
	userfirstName := getUserInput("Enter Your Firstname.!!!")
	userlastName := getUserInput("Enter Your Lastname!!!")
	userbirthday := getUserInput("Enter Your Date Of Birth DD/MM/YYYY!!!")

	var Appuser User
	Appuser = User{
		firstName: userfirstName,
		lastName:  userlastName,
		birthday:  userbirthday,
		createdAt: time.Now(),
	}

	Output(&Appuser)

	Appuser.OutputStruct();

}
// get Value of Struct and the copy og user created in memory
// func Output(user User) {
// 	fmt.Println(user.firstName, user.lastName, user.birthday,user.createdAt)
// }
func Output(user *User) {
	fmt.Println(user.firstName, user.lastName, user.birthday,user.createdAt)
	// use de refernce
	fmt.Println((*user).firstName, (*user).lastName, (*user).birthday,(*user).createdAt)

}
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value

}
