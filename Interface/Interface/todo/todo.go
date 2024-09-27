package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"title"`
}

func (todo Todo) Display() {
	fmt.Printf("Your Text: %s\n", todo.Text)
}

func New(Text string) *Todo {
	return &Todo{Text: Text}
}

func (todo Todo) SaveToFile() error {
	filename := strings.ReplaceAll(todo.Text, " ", "_")
	filename = strings.ToLower(filename)
	filename = filename + ".json"
	todoJson, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return os.WriteFile(filename, todoJson, 0644)

}
