package main

import "fmt"
type bot interface {
	/* interfaces are not a generic type
		Interfaces are implicit
	*/
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main(){
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting2(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	//Very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}