package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["red"] = "#FF0000"
	// colors["green"] = "#00FF00"
	// delete(colors, "red")
	// fmt.Printf("%+v", colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "is", hex)
	}
}
