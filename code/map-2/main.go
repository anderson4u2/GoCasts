package main

import "fmt"

func main() {
	// // Keys of type string, values of type string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// These 2 are equivalent ways of declaring an empty var
	// var colors map[string]string
	colors := make(map[string]string)

	colors["white"] = "#ffff"
	// colors.white is not a valid syntax, like it was with structs

	// delete the key "white"
	delete(colors, "white")

	fmt.Println(colors)
}
