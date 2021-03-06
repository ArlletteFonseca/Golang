package main
import("fmt"
		"net/http"		
) 

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	//a channel is like text messaging a two way channel device
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c )
	}
	fmt.Println(<-c)
}

//make get request
func checkLink(link string, c chan string) {
	//we dont' care about response here
	_,err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "yup it's up"
}