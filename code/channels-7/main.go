package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Creating a channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

// For the channel we'll have to say AGAIN the type of data that will be shared
func checkLink(link string, c chan string) {
	// Allegedly we don't need the body to verify wether the website is up
	_, err := http.Get(link)

	// it must be something's wrong with the webiste
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "MIGHT BE DOWN I THINK"
		return
	}

	fmt.Println(link, "is up!")
	c <- "YES THE SIDE LOOKS UP I THINK"
}
