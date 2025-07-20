package main

import (
	"fmt"
	"net/http"
	"os"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/ditto"

func main() {
	requestURL := pokeApiUrl
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

}
