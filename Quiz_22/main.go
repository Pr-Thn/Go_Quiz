package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"https://www.wikipedia.org/",
		"https://www.google.co.th/maps/@18.3170581,99.3986862,17z?hl=th",
		"https://pkg.go.dev/",
	}

	for _, website := range websites {
		if isWebsiteUp(website) {
			fmt.Printf("Website %s is up\n", website)
		} else {
			fmt.Printf("Website %s is down\n", website)
		}
	}
}

func isWebsiteUp(website string) bool {
	resp, err := http.Get(website)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
