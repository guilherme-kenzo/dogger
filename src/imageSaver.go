package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

// DownloadToFile downloads a file to a given filepath
// It writes file gradually, as it downloads.
func DownloadToFile(url string) (string, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := ioutil.TempFile("", "dogger_*.jpg")

	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	filename := out.Name()
	return filename, err
}
