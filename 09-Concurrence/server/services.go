package main

import (
	"fmt"
	"net/http"
)

func check(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println(server, "Error")
	} else {
		fmt.Println(server, "OK")
	}
}

func main() {
	servers := []string{"https://www.youtube.com", "https://www.google.com"}
	for _, server := range servers {
		check(server)
	}
}
