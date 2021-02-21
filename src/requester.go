package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//DogResponse is the struct for the API response from the dog.ceo api.
type DogResponse struct {
	Message string
	Status  string
}

//GetRandomDog makes a request to the https://dog.ceo/api/breeds/image/random
//endpoint and returns the response.
func GetRandomDog() DogResponse {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var d DogResponse
	json.Unmarshal(body, &d)
	return d
}
