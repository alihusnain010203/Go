package main

import (
	"errors"
	"fmt"
)

func main() {

}

func getUserInput(title string) (value string, err error) {
	if title == "" {
		err = errors.New("title is must be not empty")
		return nil, err
	}
	fmt.Println(title)

	return
}
