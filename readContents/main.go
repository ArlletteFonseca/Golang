package main

import (
	"fmt"
	"os"
	"io"	
)


	//one example to read text file and print at terminal
	/* 
		func main() {
		content, err := os.ReadFile("notes.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	}
	*/

	func main() {
		// os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, f)
	}
