package main

import (
	"fmt"
	"net/http"
)

func check(server string, chanel chan string) {
	_, err := http.Get(server)

	if err != nil {
		chanel <- server + " - Error"
	} else {
		chanel <- server + " - Ok"
	}
}

func main() {
	servers := []string{"https://www.youtube.com", "https://www.google.com"}
	chanel := make(chan string)

	for _, server := range servers {
		go check(server, chanel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-chanel)
	}
}
