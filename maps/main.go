package main

import "fmt"

func main() {
	//USE MAPS IF YOU DON"T KNOW WHAT YOUR KEYS WILL BE AT COMPILE TIME
	//keys have to be same type and values have to be same type
	//here declaring a map where all the keys are type string and all values are of type string
	colors := map[string] string {
		"red":"#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	 } 
	//below you are not assigning a value so go will print out empty map[]
	// var colors map[string]string
	//same as above
	/* colors :=make(map[string]string) */
	//use square brace syntax with map
	/* colors["white"]= "#ffffff" */
	//delete white
	/* delete(colors, "white") */
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for",color, "is", hex)
	}
}