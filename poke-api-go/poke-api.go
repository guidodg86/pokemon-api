package main

import (
	"fmt"
	"io"
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

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

}
