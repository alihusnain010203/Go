package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/note"
	"example.com/notepad/todo"
	// "github.com/brianvoe/gofakeit/v7/data"
)

type saver interface {
	SaveToFile() error
}

// Embedding the saver interface in the output interface
type output interface {
	saver
	Display()
}

func main() {
	title, desc := getNote()
	TodoText := getTodo()
	result := add(1, 2)
	println(result)

	userTodo := todo.New(TodoText)
	Output(userTodo)
	// userTodo.Display()
	// SaveData(userTodo)

	userNote := note.New(title, desc)
	Output(userNote)

	// userNote.Display()
	// SaveData(userNote)

	PrintSomething(1)
	PrintSomething(1.0)
	PrintSomething("Hello")
	PrintSomething(userTodo)
	PrintSomething(userNote)

}

func PrintSomething(value interface{}) {
	// switch value.(type) {
	// 	case todo.Todo:
	// 		fmt.Println("This is a Todo")
	// 	case note.Note:
	// 		fmt.Println("This is a Note")
	// 	case int:
	// 		fmt.Println("This is an integer")
	// 	case float64:
	// 		fmt.Println("This is a float")
	// 	case string:
	// 		fmt.Println("This is a string")
	// 	default:
	// 		fmt.Println("Unknown type")
	// }


	// 2nd way
	Type, ok := value.(todo.Todo)
	fmt.Println(Type, ok)


}

func Output(data output) {
	data.Display()
	data.SaveToFile()

}
func getTodo() string {
	Text, err := getUserInput("Enter Text...")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return Text
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
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input, nil
}

func SaveData(data saver) error {
	data.SaveToFile()
	return nil
}

func add[T int|float64|string|interface{}	](a,b T)T{
	return a + b
}