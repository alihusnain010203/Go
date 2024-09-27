package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/note"
)

func main() {
	title,desc:=getNote()
	
	userNote:=note.New(title,desc)
	userNote.Display()
	userNote.SaveToFile()
}
func getNote() (string, string) {
	title, err := getUserInput("Enter Title...")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	content, err := getUserInput("Enter Description ...")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return title, content 
}
func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// Remove the newline character from the end of the input
	input = strings.TrimSuffix(input,"\n")
	input=strings.TrimSuffix(input,"\r")
	return input, nil
}


