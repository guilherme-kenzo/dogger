package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	randomDog := GetRandomDog()
	resp, err := http.Get(string(randomDog.Message))

	if err != nil {
		log.Fatal(err)
	}

	filename, err = DownloadToFile()

	asciistring := ConvertImageToASCII(filename)

	fmt.Printf(asciistring)

}
