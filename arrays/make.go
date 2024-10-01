package main

import "fmt"

func main() {
	// userName := make([]string, 2, 5)
	// userName = append(userName, "John")
	// userName = append(userName, "Doe")
	// userName = append(userName, "Smith")
	// userName = append(userName, "Jane")
	// userName= append(userName, "Doe")
	// userName= append(userName, "Smith")
	// fmt.Println(userName)
	webSites := make(map[string]string,3)
	webSites["Google"] = "https://www.google.com"
	webSites["AWS"]="https://aws.amazon.com"
	webSites["Facebook"] = "https://www.facebook.com"
	// linkedIn
	webSites["LinkedIn"] = "https://www.linkedin.com"
	link := webSites["Google"]
	fmt.Println(link)

	// iterate over the map
	for key, value := range webSites {
		fmt.Println(key, value)
	}
	// delete an element
	delete(webSites, "AWS")

}