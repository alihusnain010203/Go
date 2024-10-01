package maps

import "fmt"

func maps() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}
	fmt.Println(websites["Google"])
	websites["AWS"]= "https://aws.amazon.com"
	fmt.Println(websites["AWS"])
	// delete an element
	delete(websites, "AWS")
	fmt.Println(websites)
	

}