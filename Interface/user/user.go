package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

// Embedding
type EmbeddedUser struct {
	user User
	age  string
}

func (user *User) OutputStruct() {
	fmt.Println(user.firstName, user.lastName, user.birthday, user.createdAt)
	// use de refernce
	fmt.Println((*user).firstName, (*user).lastName, (*user).birthday, (*user).createdAt)
}

//	func (user User) Clear(){
//		user.firstName = ""
//		user.lastName=""
//	}
func (user *User) Clear() {
	user.firstName = ""
	user.lastName = ""
}

// Constructor function

func New(fn, ln, bd string) (*User, error) {
	// Validation
	if fn == "" || ln == "" || bd == "" {
		//   error
		return nil, errors.New("all field required")
	}
	return &User{
		firstName: fn,
		lastName:  ln,
		birthday:  bd,
		createdAt: time.Now(),
	}, nil
}
