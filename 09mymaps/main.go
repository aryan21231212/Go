package main

import "fmt"

func main() {

	fmt.Println("Maps in Go!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Go"

	fmt.Println("List of languages: ", languages)
	fmt.Println("JS stands for: ", languages["JS"])

	// Delete an element
	delete(languages, "GO")
	fmt.Println("List of languages: ", languages)

	// Loop through map
	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}
}
