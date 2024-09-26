package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n Note) Display() {
	fmt.Printf("Your Title: %s\n", n.Title)
	fmt.Printf("Your Content: %s\n", n.Content)
}

func New(Title, Content string) *Note {
	return &Note{Title, Content}
}

func (note Note) SaveToFile() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename)
	filename = filename + ".json"
	noteJson, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return os.WriteFile(filename, noteJson, 0644)

}
