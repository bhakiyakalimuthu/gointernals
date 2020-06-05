package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal("failed to call server from client")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	io.Copy(os.Stdout, res.Body)

}