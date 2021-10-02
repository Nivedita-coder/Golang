package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("List of all languages: ", languages)

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)
}
