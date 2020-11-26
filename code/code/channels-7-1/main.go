package main

import (
	"fmt"
	"net/http"
	"time"
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

	// INFINITE LOOP
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative syntax to INFINITE LOOP
	// l short for link
	for l := range c {
		// go checkLink(l, c)
		// Replacing above with a function literal
		// `l` shouldn't be used here, a child routine..
		// SHOULD NOT reuse a var from the main routine
		// Only pass as argument or talk through Channel
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// For the channel we'll have to say AGAIN the type of data that will be shared
func checkLink(link string, c chan string) {
	// Allegedly we don't need the body to verify wether the website is up
	_, err := http.Get(link)

	// it must be something's wrong with the webiste
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
