package main

import "fmt"

func main() {
	// colors := map[string]string {
	// 	"red":"#ff000",
	// }

	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	fmt.Println(colors)
}