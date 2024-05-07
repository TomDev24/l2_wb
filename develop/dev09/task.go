package main

import (
	"log"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Specify URL")
	}

	url := os.Args[1]
	filename := filepath.Base(url)

	err := download(url, filename)
	if err != nil {
		log.Fatal(err)
	}
}
