package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
func(logWriter) Write(bs []byte)(int, error){
	fmt.Println(string(bs))
	fmt.Println("bytes", len(bs))
	return len(bs), nil
}