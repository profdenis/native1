package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadToFile(url string, filepath string) error {
	// Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create the file
	outFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write the body to file
	_, err = io.Copy(outFile, response.Body)
	return err
}

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls, scanner.Err()
}

func downloadAllWithGoroutines(urls []string) {
	ch := make(chan bool)
	for i, url := range urls {
		go func(address string, j int) {
			err := downloadToFile(address, fmt.Sprintf("fichiers/downloads/address-%d.html", j))
			if err != nil {
				fmt.Println(err)
				ch <- false
			} else {
				ch <- true
			}

		}(url, i)
	}
	for i := 0; i < len(urls); i++ {
		<-ch
	}
}

func downloadAllWithoutGoroutines(urls []string, err error) {
	for i, url := range urls {
		err = downloadToFile(url, fmt.Sprintf("fichiers/downloads/url-%d.html", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
