package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// url defines the endpoint link for the file to be downloaded
	url := os.Args[1]

	if len(url) == 0 {
		log.Fatal("No provided url link\n The format should be `vuta https://github.com/jkarage/vuta.git`")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Content-Length: ", resp.ContentLength)
}
